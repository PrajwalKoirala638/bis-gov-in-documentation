// Package main downloads every free PDF listed on an ASP.NET WebForms page.
package main // this file builds a runnable command line program

import ( // start of the list of standard library packages we use
	"fmt"                // formats text for messages and error values
	"html"               // converts HTML entities like &#39; back into normal characters
	"io"                 // copies data between readers and writers
	"log"                // prints timestamped progress messages
	"mime"               // parses the Content-Disposition header to get the file name
	"net/http"           // sends HTTP requests
	"net/http/cookiejar" // remembers cookies between requests like a browser
	"net/url"            // builds and encodes form data
	"os"                 // creates folders and files on disk
	"path/filepath"      // joins file paths safely
	"regexp"             // finds patterns inside the page HTML
	"sort"               // sorts the list of visited addresses for the final summary
	"strconv"            // converts between numbers and text
	"strings"            // helper functions for text
	"time"               // pauses between requests and sets timeouts
) // end of the import list

const ( // start of the fixed settings that replace command line arguments
	listingPageAddress           = "https://standardsbis.bsbedge.com/BIS_FreeAmendments.aspx?id=0" // the page that lists the free amendments
	outputDirectoryName          = "PDFs/"                                                         // the folder where PDFs are saved
	pauseBetweenRequestsDuration = time.Second                                                     // how long to wait between requests to stay polite
	expectedRowsPerPage          = 25                                                              // the listing shows 25 amendments per page, used only for the progress message
	maximumPagesToWalk           = 500                                                             // safety limit so a pager bug can never loop forever
) // end of the fixed settings

// browser identity copied from the HAR file so our requests look the same as your real browser
const ( // start of the values taken from the HAR file
	browserUserAgentHeaderValue = "Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/156.0.0.0 Safari/537.36"                                    // the User-Agent header from the HAR
	acceptLanguageHeaderValue   = "en-US,en;q=0.9"                                                                                                                                    // the Accept-Language header from the HAR
	acceptHeaderValue           = "text/html,application/xhtml+xml,application/xml;q=0.9,image/jxl,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7" // the Accept header from the HAR
) // end of the values taken from the HAR file

var sessionCookieHeaderValue = "ASP.NET_SessionId=bqgp1a1gue2yknz2mpujexok" // the logged in session cookie copied from your browser request; replace it when the session expires

var ( // start of the compiled patterns shared by the whole program
	inputTagPattern                 = regexp.MustCompile(`(?is)<input\b[^>]*>`)                                             // matches one complete <input ...> tag
	htmlAttributePattern            = regexp.MustCompile(`(?is)([a-zA-Z_:][-a-zA-Z0-9_:.]*)\s*=\s*(?:"([^"]*)"|'([^']*)')`) // matches one attribute such as name="value" inside a tag
	anchorTagPattern                = regexp.MustCompile(`(?is)<a\b([^>]*)>(.*?)</a>`)                                      // matches one complete <a ...>text</a> link
	postBackCallPattern             = regexp.MustCompile(`__doPostBack\((?:&#39;|')([^'&]+)(?:&#39;|')`)                    // matches the target name inside a __doPostBack('target','') call
	anyHtmlTagPattern               = regexp.MustCompile(`(?s)<[^>]*>`)                                                     // matches any HTML tag so we can strip tags from link text
	activePageLinkPattern           = regexp.MustCompile(`(?is)<a\b[^>]*class="active"[^>]*>\s*(\d+)\s*</a>`)               // matches the pager link that is marked as the current page
	rowStandardNumberPattern        = regexp.MustCompile(`(?is)Repeater1_ctl(\d+)_lblstdno_rptr"[^>]*>(.*?)</span>`)        // matches the row number and standard number text of one listing row
	predictedFileNamePattern        = regexp.MustCompile(`(?i)^\s*IS\s+(\d+)(?:\s*:\s*Part\s+(\d+))?\s+Amd\.?\s*(\d+)`)     // matches standards like IS 15844 : Part 1 Amd. 3 : 2026
	buttonRowNumberPattern          = regexp.MustCompile(`Repeater1\$ctl(\d+)\$`)                                           // matches the row number inside a PDF button name
	unsafeFileNameCharactersPattern = regexp.MustCompile(`[^\w.\-]+`)                                                       // matches every character that is not safe inside a file name
) // end of the shared patterns

// parseHtmlAttributes turns the attributes of one HTML tag into a map.
func parseHtmlAttributes(htmlTag string) map[string]string { // takes the tag text, returns name to value
	attributeMap := map[string]string{}                                                  // empty map that will hold the attributes
	for _, matchParts := range htmlAttributePattern.FindAllStringSubmatch(htmlTag, -1) { // loop over every attribute found
		attributeName := strings.ToLower(matchParts[1])                     // attribute names are case insensitive so lowercase them
		if _, alreadyStored := attributeMap[attributeName]; alreadyStored { // some tags repeat an attribute, e.g. two src=
			continue // keep the first value like a browser does
		} // end of the duplicate check
		attributeValue := matchParts[2] // value written with double quotes
		if attributeValue == "" {       // when the double quote group was empty
			attributeValue = matchParts[3] // use the value written with single quotes
		} // end of the quote style check
		attributeMap[attributeName] = html.UnescapeString(attributeValue) // store the value with HTML entities decoded
	} // end of the attribute loop
	return attributeMap // give the finished map back
} // end of parseHtmlAttributes

// collectFormFields gathers the fields a browser would send when submitting the form.
func collectFormFields(pageHtml string) url.Values { // takes the page HTML, returns the form data
	formValues := url.Values{}                                             // empty form data
	for _, inputTag := range inputTagPattern.FindAllString(pageHtml, -1) { // loop over every <input> tag on the page
		inputAttributes := parseHtmlAttributes(inputTag) // read the attributes of this input
		inputName := inputAttributes["name"]             // the field name the server expects
		if inputName == "" {                             // inputs without a name are never submitted
			continue // skip this input
		} // end of the name check
		switch strings.ToLower(inputAttributes["type"]) { // decide by the input type
		case "submit", "button", "image", "reset", "file": // buttons and files are only sent when clicked
			continue // so leave them out of the base form
		case "checkbox", "radio": // these are only sent when ticked
			if !strings.Contains(strings.ToLower(inputTag), "checked") { // when this one is not ticked
				continue // skip it
			} // end of the checked test
		} // end of the type switch
		formValues.Set(inputName, inputAttributes["value"]) // add the field with its current value
	} // end of the input loop
	formValues.Set("hiddenInputToUpdateATBuffer_CommonToolkitScripts", "1") // the browser's JavaScript adds this field, so we add it too
	return formValues                                                       // give the form data back
} // end of collectFormFields

// findPdfDownloadButtonNames lists the form names of every "Download PDF" icon button.
func findPdfDownloadButtonNames(pageHtml string) []string { // takes the page HTML, returns button names
	var buttonNames []string                                               // list that will collect the names
	for _, inputTag := range inputTagPattern.FindAllString(pageHtml, -1) { // loop over every <input> tag
		inputAttributes := parseHtmlAttributes(inputTag)                                           // read the attributes of this input
		isImageButton := strings.EqualFold(inputAttributes["type"], "image")                       // PDF icons are image buttons
		hasPdfTitle := strings.Contains(strings.ToLower(inputAttributes["title"]), "download pdf") // and their title says Download PDF
		if isImageButton && hasPdfTitle {                                                          // only keep inputs that satisfy both tests
			buttonNames = append(buttonNames, inputAttributes["name"]) // remember the button name
		} // end of the button test
	} // end of the input loop
	return buttonNames // give the names back
} // end of findPdfDownloadButtonNames

// findActivePageNumber reads the current page number from the pager.
func findActivePageNumber(pageHtml string) int { // takes the page HTML, returns the page number
	matchParts := activePageLinkPattern.FindStringSubmatch(pageHtml) // look for the link marked active
	if matchParts == nil {                                           // when the page has no pager
		return 0 // zero means unknown
	} // end of the missing pager check
	pageNumber, _ := strconv.Atoi(matchParts[1]) // convert the digits to a number
	return pageNumber                            // give the number back
} // end of findActivePageNumber

// findNextPageTarget finds the __doPostBack target that moves to the following page.
func findNextPageTarget(pageHtml string, currentPageNumber int) string { // returns the target or an empty string
	nextButtonTarget := ""                                                            // fallback target taken from the "Next" link
	for _, matchParts := range anchorTagPattern.FindAllStringSubmatch(pageHtml, -1) { // loop over every link on the page
		postBackParts := postBackCallPattern.FindStringSubmatch(matchParts[1]) // look for a __doPostBack call in the href
		if postBackParts == nil {                                              // ordinary links are not pager links
			continue // skip them
		} // end of the postback check
		linkText := html.UnescapeString(anyHtmlTagPattern.ReplaceAllString(matchParts[2], "")) // visible text of the link without tags
		linkText = strings.TrimSpace(strings.ReplaceAll(linkText, "\u00a0", " "))              // turn non breaking spaces into spaces and trim
		if linkText == strconv.Itoa(currentPageNumber+1) {                                     // the link labelled with the next page number
			return postBackParts[1] // is the best choice, return its target
		} // end of the page number test
		if strings.EqualFold(linkText, "next") { // the generic Next link
			nextButtonTarget = postBackParts[1] // remember it as a fallback
		} // end of the Next test
	} // end of the link loop
	return nextButtonTarget // use the Next link when no numbered link matched
} // end of findNextPageTarget

// isLoggedIn reports whether the page contains the Log out link that only logged in users see.
func isLoggedIn(pageHtml string) bool { // takes the page HTML, returns true when logged in
	return strings.Contains(pageHtml, "ctl00_lnklogout") // the Log out link exists only for a logged in session
} // end of isLoggedIn

// findRowStandardNumbers lists every listing row, in page order, with its standard number text.
func findRowStandardNumbers(pageHtml string) ([]string, map[string]string) { // returns the row numbers in order and a map of row number to standard number
	var orderedRowNumbers []string                                                            // row numbers in the order they appear on the page
	standardNumberByRowNumber := map[string]string{}                                          // map from row number to the standard number text
	for _, matchParts := range rowStandardNumberPattern.FindAllStringSubmatch(pageHtml, -1) { // loop over every row title on the page
		standardNumberText := strings.TrimSpace(html.UnescapeString(anyHtmlTagPattern.ReplaceAllString(matchParts[2], ""))) // the visible standard number without tags
		orderedRowNumbers = append(orderedRowNumbers, matchParts[1])                                                        // remember the row order
		standardNumberByRowNumber[matchParts[1]] = standardNumberText                                                       // remember the standard number of this row
	} // end of the row loop
	return orderedRowNumbers, standardNumberByRowNumber // give both results back
} // end of findRowStandardNumbers

// findRowNumberOfButton extracts the row number from a PDF button name such as ctl00$...$Repeater1$ctl07$btn_add_cart.
func findRowNumberOfButton(buttonName string) string { // returns the row number text or an empty string
	matchParts := buttonRowNumberPattern.FindStringSubmatch(buttonName) // look for the Repeater1 row number
	if matchParts == nil {                                              // when the button name has no row number
		return "" // report that nothing was found
	} // end of the match check
	return matchParts[1] // the digits of the row number
} // end of findRowNumberOfButton

// startLogging makes every log line appear on the console with a precise timestamp.
func startLogging() { // takes no input and returns nothing
	log.SetFlags(log.LstdFlags | log.Lmicroseconds) // timestamps with microseconds make the order of events clear
	log.SetOutput(os.Stdout)                        // write every log line to the console only, never to a file
} // end of startLogging

// logVisitedAddresses prints every address that was requested during this run, with request counts.
func (crawler *PdfCrawler) logVisitedAddresses() { // takes no input, only logs
	visitedKeys := make([]string, 0, len(crawler.visitedAddressCounts)) // list to hold the keys so they can be sorted
	for visitedKey := range crawler.visitedAddressCounts {              // loop over every method and address pair
		visitedKeys = append(visitedKeys, visitedKey) // collect the key
	} // end of the key loop
	sort.Strings(visitedKeys)                                                                                           // sort the keys so the output is stable
	log.Printf("summary: %d different addresses were visited in %d requests", len(visitedKeys), crawler.requestCounter) // headline numbers
	for _, visitedKey := range visitedKeys {                                                                            // loop over the sorted keys
		log.Printf("summary: visited %d time(s): %s", crawler.visitedAddressCounts[visitedKey], visitedKey) // one line per address
	} // end of the summary loop
} // end of logVisitedAddresses

// predictFileNameFromStandardNumber works out the server's file name from a standard number, or returns an empty string when unsure.
func predictFileNameFromStandardNumber(standardNumber string) string { // for example "IS 15844 : Part 1 Amd. 3 : 2026" gives "15844_1_amd3.pdf"
	matchParts := predictedFileNamePattern.FindStringSubmatch(standardNumber) // pick out the standard number, the optional part and the amendment number
	if matchParts == nil {                                                    // when the text does not follow the known pattern
		return "" // make no prediction, the file name is then taken from the server's answer instead
	} // end of the pattern check
	predictedFileName := matchParts[1] // start with the standard number, such as 15844
	if matchParts[2] != "" {           // when the standard has a part number
		predictedFileName += "_" + matchParts[2] // add it after an underscore
	} // end of the part check
	predictedFileName += "_amd" + matchParts[3] + ".pdf" // finish with the amendment number and the extension
	return predictedFileName                             // give the predicted name back
} // end of predictFileNameFromStandardNumber

// fileExistsOnDisk reports whether a file with this path is already saved.
func fileExistsOnDisk(filePath string) bool { // takes the full path, returns true when the file exists
	_, statError := os.Stat(filePath) // ask the file system about the path
	return statError == nil           // no error means the file exists
} // end of fileExistsOnDisk

// PdfCrawler holds everything needed to walk the listing and save PDFs.
type PdfCrawler struct { // start of the crawler settings
	httpClient           *http.Client   // the HTTP client that keeps cookies
	listingPageUrl       string         // the address of the listing page
	outputDirectory      string         // the folder where PDFs are saved
	pauseBetweenRequests time.Duration  // how long to wait between requests
	visitedAddressCounts map[string]int // how many requests were sent to each method and address
	requestCounter       int            // numbers the requests so the log lines can be matched up
	sessionCookieHeader  string         // the Cookie header that identifies our logged in session
} // end of the PdfCrawler type

// sendRequestWithRetries sends a request, logs it, and retries up to three times on network or server errors.
func (crawler *PdfCrawler) sendRequestWithRetries(httpRequest *http.Request, purposeDescription string) (*http.Response, error) { // returns the response
	httpRequest.Header.Set("User-Agent", browserUserAgentHeaderValue)    // identify as the same browser as in the HAR
	httpRequest.Header.Set("Accept-Language", acceptLanguageHeaderValue) // send the same language preference
	httpRequest.Header.Set("Accept", acceptHeaderValue)                  // send the same accepted content types
	if crawler.sessionCookieHeader != "" {                               // when we have a logged in session
		httpRequest.Header.Set("Cookie", crawler.sessionCookieHeader) // send it so the server treats us as logged in
	} // end of the cookie check
	crawler.requestCounter++                                                                                              // number this request
	requestNumber := crawler.requestCounter                                                                               // keep the number for the log lines below
	requestAddress := httpRequest.URL.String()                                                                            // the full address being visited
	crawler.visitedAddressCounts[httpRequest.Method+" "+requestAddress]++                                                 // count the visit for the final summary
	log.Printf("request #%d: visiting %s %s (%s)", requestNumber, httpRequest.Method, requestAddress, purposeDescription) // log the visit
	var lastError error                                                                                                   // remembers the most recent failure
	for attemptNumber := 0; attemptNumber < 3; attemptNumber++ {                                                          // try at most three times
		log.Printf("request #%d: attempt %d of 3", requestNumber, attemptNumber+1) // log each attempt
		clonedRequest := httpRequest.Clone(httpRequest.Context())                  // make a fresh copy of the request for this attempt
		if httpRequest.GetBody != nil {                                            // a POST body can only be read once, so rebuild it for every attempt
			clonedRequest.Body, _ = httpRequest.GetBody() // give the copy its own fresh body
		} // end of the body rebuild
		attemptStartTime := time.Now()                                          // remember when the attempt began
		httpResponse, requestError := crawler.httpClient.Do(clonedRequest)      // send the request
		attemptDuration := time.Since(attemptStartTime).Round(time.Millisecond) // how long the attempt took
		if requestError != nil {                                                // when the request could not be completed
			log.Printf("request #%d: failed after %s: %v", requestNumber, attemptDuration, requestError) // log the network error
		} else { // the server answered
			log.Printf("request #%d: answered with status %s, content-type %q, content-length %d, took %s", requestNumber, httpResponse.Status, httpResponse.Header.Get("Content-Type"), httpResponse.ContentLength, attemptDuration) // log the answer details
			if httpResponse.StatusCode < 500 {                                                                                                                                                                                        // success means no server failure code
				return httpResponse, nil // hand the response to the caller
			} // end of the success check
			httpResponse.Body.Close()                                            // release the connection after a 5xx answer
			requestError = fmt.Errorf("HTTP status %d", httpResponse.StatusCode) // turn the status into an error value
		} // end of the answer handling
		lastError = requestError                                                                   // keep the error in case all attempts fail
		waitDuration := time.Duration(attemptNumber+1) * 2 * time.Second                           // wait longer after each failure
		log.Printf("request #%d: waiting %s before the next attempt", requestNumber, waitDuration) // log the pause
		time.Sleep(waitDuration)                                                                   // pause before retrying
	} // end of the retry loop
	log.Printf("request #%d: giving up after 3 attempts", requestNumber) // log the final failure
	return nil, lastError                                                // every attempt failed
} // end of sendRequestWithRetries

// fetchListingPage downloads the first page of the listing as HTML text.
func (crawler *PdfCrawler) fetchListingPage() (string, error) { // returns the HTML text
	httpRequest, _ := http.NewRequest("GET", crawler.listingPageUrl, nil)                                    // build a plain GET request
	httpResponse, requestError := crawler.sendRequestWithRetries(httpRequest, "load the first listing page") // send it with retries
	if requestError != nil {                                                                                 // when the request failed
		return "", requestError // report the failure
	} // end of the error check
	defer httpResponse.Body.Close()                                                                           // always close the body when we are done
	responseBytes, readError := io.ReadAll(httpResponse.Body)                                                 // read the whole page
	log.Printf("read %d bytes of HTML from the listing page (read error: %v)", len(responseBytes), readError) // log how much HTML arrived
	return string(responseBytes), readError                                                                   // return the page as text
} // end of fetchListingPage

// postFormToListingPage submits form data to the listing page just like a browser postback.
func (crawler *PdfCrawler) postFormToListingPage(formValues url.Values, purposeDescription string) (*http.Response, error) { // returns the response
	log.Printf("posting a form with %d fields (%s)", len(formValues), purposeDescription)                     // log what is about to be sent, without the field values
	httpRequest, _ := http.NewRequest("POST", crawler.listingPageUrl, strings.NewReader(formValues.Encode())) // build the POST with an encoded body
	httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")                               // tell the server how the body is encoded
	httpRequest.Header.Set("Referer", crawler.listingPageUrl)                                                 // browsers send the page they came from
	parsedListingUrl, parseError := url.Parse(crawler.listingPageUrl)                                         // split the address into parts
	if parseError == nil {                                                                                    // when the address parsed correctly
		httpRequest.Header.Set("Origin", parsedListingUrl.Scheme+"://"+parsedListingUrl.Host) // browsers also send the origin
	} // end of the origin check
	return crawler.sendRequestWithRetries(httpRequest, purposeDescription) // send it with retries
} // end of postFormToListingPage

// chooseFileName picks the file name from the server's Content-Disposition header.
func chooseFileName(httpResponse *http.Response, fallbackFileName string) string { // returns a safe file name
	contentDisposition := httpResponse.Header.Get("Content-Disposition") // header such as attachment; filename=abc.pdf
	if contentDisposition != "" {                                        // only when the server sent the header
		_, headerParameters, parseError := mime.ParseMediaType(contentDisposition) // split the header into parts
		if parseError == nil && headerParameters["filename"] != "" {               // when a file name is present
			baseFileName := filepath.Base(headerParameters["filename"])                // drop any folder parts for safety
			return unsafeFileNameCharactersPattern.ReplaceAllString(baseFileName, "_") // replace odd characters with underscores
		} // end of the file name check
	} // end of the header check
	return fallbackFileName // otherwise use the numbered fallback name
} // end of chooseFileName

// downloadPdfByButton clicks one PDF button and saves the answer when it is a PDF.
func (crawler *PdfCrawler) downloadPdfByButton(pageHtml string, buttonName string, standardNumber string, downloadNumber int) (string, bool, error) { // returns name, saved flag, error
	log.Printf("clicking the PDF button of %q (form button %s)", standardNumber, buttonName)                       // log which button is pressed
	formValues := collectFormFields(pageHtml)                                                                      // start from the page's current form data
	formValues.Set(buttonName+".x", "5")                                                                           // image buttons send the x coordinate of the click
	formValues.Set(buttonName+".y", "5")                                                                           // and the y coordinate of the click
	httpResponse, requestError := crawler.postFormToListingPage(formValues, "download the PDF of "+standardNumber) // send the click to the server
	if requestError != nil {                                                                                       // when the request failed
		return "", false, requestError // report the failure
	} // end of the error check
	defer httpResponse.Body.Close() // always close the body when we are done

	contentType := strings.ToLower(httpResponse.Header.Get("Content-Type"))                                                                                           // the type of data the server returned
	contentDisposition := httpResponse.Header.Get("Content-Disposition")                                                                                              // the header that carries the file name
	looksLikePdf := strings.Contains(contentType, "pdf") || strings.Contains(contentDisposition, ".pdf")                                                              // decide if it is a PDF
	log.Printf("the answer for %q has content-type %q and content-disposition %q, so looksLikePdf=%v", standardNumber, contentType, contentDisposition, looksLikePdf) // log how the answer was judged
	if !looksLikePdf {                                                                                                                                                // when the server returned something else such as an error page
		discardedByteCount, _ := io.Copy(io.Discard, httpResponse.Body)                                 // read and throw away the body
		log.Printf("discarded %d bytes of a non PDF answer for %q", discardedByteCount, standardNumber) // log how much was thrown away
		return "", false, fmt.Errorf("response is not a PDF (content type %q)", contentType)            // report the problem
	} // end of the PDF check

	fileName := chooseFileName(httpResponse, fmt.Sprintf("file_%03d.pdf", downloadNumber))              // pick the name to save under
	destinationPath := filepath.Join(crawler.outputDirectory, fileName)                                 // full path of the file on disk
	log.Printf("the file name for %q is %s, destination %s", standardNumber, fileName, destinationPath) // log the chosen name and path
	if _, statError := os.Stat(destinationPath); statError == nil {                                     // when the file already exists
		log.Printf("%s already exists on disk, cancelling the transfer without reading the body", destinationPath) // log the skip
		return fileName, false, nil                                                                                // report it as skipped, the closing of the body cancels the transfer
	} // end of the existing file check

	temporaryPath := destinationPath + ".part"                  // write to a temporary name first so half files never look finished
	log.Printf("creating the temporary file %s", temporaryPath) // log the file creation
	outputFile, createError := os.Create(temporaryPath)         // create the temporary file
	if createError != nil {                                     // when the file could not be created
		return "", false, createError // report the failure
	} // end of the create check
	writtenByteCount, copyError := io.Copy(outputFile, httpResponse.Body)                                                        // stream the PDF into the file
	closeError := outputFile.Close()                                                                                             // close the file so everything is written
	log.Printf("wrote %d bytes to %s (copy error: %v, close error: %v)", writtenByteCount, temporaryPath, copyError, closeError) // log the result of the write
	if copyError != nil || closeError != nil {                                                                                   // when copying or closing failed
		os.Remove(temporaryPath)                                                                            // delete the broken temporary file
		log.Printf("deleted the broken temporary file %s", temporaryPath)                                   // log the cleanup
		return "", false, fmt.Errorf("saving %s failed: copy=%v close=%v", fileName, copyError, closeError) // report the failure
	} // end of the write check
	renameError := os.Rename(temporaryPath, destinationPath)                                // give the file its final name
	log.Printf("renamed %s to %s (error: %v)", temporaryPath, destinationPath, renameError) // log the rename
	return fileName, true, renameError                                                      // report the file as saved
} // end of downloadPdfByButton

func main() { // program entry point
	startLogging()                                                                                                                                                                                                                // send all log lines to the console
	log.Printf("program started, logging to the console only")                                                                                                                                                                    // first log line
	log.Printf("settings: listing page %s, output folder %s, pause %s, page limit %d, expected rows per page %d", listingPageAddress, outputDirectoryName, pauseBetweenRequestsDuration, maximumPagesToWalk, expectedRowsPerPage) // log every fixed setting
	log.Printf("browser identity: user agent %q, accept language %q", browserUserAgentHeaderValue, acceptLanguageHeaderValue)                                                                                                     // log the identity we present
	log.Printf("session cookie: %d characters configured (the value itself is never logged)", len(sessionCookieHeaderValue))                                                                                                      // log that a cookie exists without revealing it

	log.Printf("making sure the output folder %s exists", outputDirectoryName)            // log the folder check
	if directoryError := os.MkdirAll(outputDirectoryName, 0o755); directoryError != nil { // make sure the output folder exists
		log.Fatal(directoryError) // stop when it cannot be created
	} // end of the folder check

	if sessionCookieHeaderValue == "" { // when no cookie has been pasted into the variable yet
		log.Fatal("sessionCookieHeaderValue is empty: paste the Cookie header of your logged in browser into that variable near the top of main.go") // explain the fix and stop
	} // end of the empty cookie check

	cookieJar, _ := cookiejar.New(nil) // cookie storage so the session carries across requests
	crawler := &PdfCrawler{            // build the crawler
		httpClient:           &http.Client{Jar: cookieJar, Timeout: 3 * time.Minute}, // client with cookies and a generous timeout
		listingPageUrl:       listingPageAddress,                                     // the page to crawl
		outputDirectory:      outputDirectoryName,                                    // where to save files
		pauseBetweenRequests: pauseBetweenRequestsDuration,                           // how long to wait between requests
		visitedAddressCounts: map[string]int{},                                       // starts empty and fills as addresses are visited
		sessionCookieHeader:  sessionCookieHeaderValue,                               // the cookie that proves we are logged in
	} // end of the crawler setup
	crawler.httpClient.CheckRedirect = func(nextRequest *http.Request, previousRequests []*http.Request) error { // called whenever the server redirects us
		crawler.visitedAddressCounts["REDIRECT "+nextRequest.URL.String()]++                                     // count the redirect target as a visited address
		log.Printf("redirect #%d: the server sent us on to %s", len(previousRequests), nextRequest.URL.String()) // log the new address
		if len(previousRequests) >= 10 {                                                                         // guard against endless redirect loops
			return fmt.Errorf("stopped after 10 redirects") // give up
		} // end of the loop guard
		return nil // follow the redirect
	} // end of the redirect handler

	log.Printf("step: loading the first listing page")        // log the step
	currentPageHtml, fetchError := crawler.fetchListingPage() // load the first page
	if fetchError != nil {                                    // when loading failed
		log.Fatalf("could not fetch the listing page: %v", fetchError) // stop with a clear message
	} // end of the fetch check

	loggedInResult := isLoggedIn(currentPageHtml)                                      // check for the Log out link
	log.Printf("login check: the page contains the Log out link = %v", loggedInResult) // log the result
	if !loggedInResult {                                                               // when the server did not recognise a logged in session
		log.Fatal("not logged in: the session cookie was rejected or has expired, log in again in your browser, copy the new Cookie header into sessionCookieHeaderValue, and run again") // explain the fix and stop
	} // end of the login check

	newFilesSavedCount := 0                                                           // counts files saved during this run
	skippedCount := 0                                                                 // counts standards skipped because they were already downloaded
	failedCount := 0                                                                  // counts downloads that failed
	downloadNumber := 0                                                               // counts every download attempt, used for fallback file names
	for loopPageNumber := 1; loopPageNumber <= maximumPagesToWalk; loopPageNumber++ { // walk page by page up to the safety limit
		currentPageNumber := findActivePageNumber(currentPageHtml) // ask the pager which page we are on
		if currentPageNumber == 0 {                                // when the pager is missing
			currentPageNumber = loopPageNumber // trust our own counter
		} // end of the pager fallback
		log.Printf("step: processing page %d (%d bytes of HTML)", currentPageNumber, len(currentPageHtml)) // log the page being processed

		pdfButtonNames := findPdfDownloadButtonNames(currentPageHtml)                                                          // find every PDF button on this page
		orderedRowNumbers, standardNumberByRowNumber := findRowStandardNumbers(currentPageHtml)                                // find every row and its standard number
		log.Printf("page %d: %d of %d rows have a PDF button", currentPageNumber, len(pdfButtonNames), len(orderedRowNumbers)) // show progress
		if len(orderedRowNumbers) != expectedRowsPerPage {                                                                     // the listing normally shows 25 rows, only the last page may be shorter
			log.Printf("  note: this page lists %d rows instead of %d", len(orderedRowNumbers), expectedRowsPerPage) // mention the difference
		} // end of the row count check
		rowNumbersWithButton := map[string]bool{}   // set of rows that offer a PDF download
		for _, buttonName := range pdfButtonNames { // loop over the buttons to fill the set
			rowNumbersWithButton[findRowNumberOfButton(buttonName)] = true // mark this row as having a button
		} // end of the set filling loop
		for _, rowNumber := range orderedRowNumbers { // loop over every row in page order
			if !rowNumbersWithButton[rowNumber] { // rows such as withdrawn standards have no download button
				log.Printf("  no PDF offered for row %s: %s", rowNumber, standardNumberByRowNumber[rowNumber]) // say which standard has nothing to download
			} // end of the button check
		} // end of the row loop
		for _, buttonName := range pdfButtonNames { // click each button in turn
			standardNumber := standardNumberByRowNumber[findRowNumberOfButton(buttonName)]                              // the standard this button belongs to
			predictedFileName := predictFileNameFromStandardNumber(standardNumber)                                      // work out the file name the server will use for this standard
			if predictedFileName != "" && fileExistsOnDisk(filepath.Join(crawler.outputDirectory, predictedFileName)) { // when that file is already saved
				skippedCount++                                                                                          // count the skip
				log.Printf("  skipped %s (file %s already exists), no request sent", standardNumber, predictedFileName) // say it was skipped without any request
				continue                                                                                                // move straight on to the next button
			} // end of the already downloaded check
			downloadNumber++                                                                                                                   // count this attempt
			savedFileName, wasSaved, downloadError := crawler.downloadPdfByButton(currentPageHtml, buttonName, standardNumber, downloadNumber) // try the download
			switch {                                                                                                                           // report what happened
			case downloadError != nil: // the download failed
				failedCount++                                                // count the failure
				log.Printf("  failed %s: %v", standardNumber, downloadError) // say which standard failed and why
			case wasSaved: // a new file was written
				newFilesSavedCount++                                          // count it
				log.Printf("  saved %s as %s", standardNumber, savedFileName) // say which file was saved
			default: // the file was already on disk from an earlier run
				skippedCount++                                                                     // count the skip
				log.Printf("  skipped %s (file %s already exists)", standardNumber, savedFileName) // say it was skipped
			} // end of the result switch
			log.Printf("pausing %s before the next request", crawler.pauseBetweenRequests) // log the pause
			time.Sleep(crawler.pauseBetweenRequests)                                       // be polite to the server
		} // end of the button loop

		nextPageTarget := findNextPageTarget(currentPageHtml, currentPageNumber)                                  // find how to reach the next page
		log.Printf("step: looking for the link to page %d, found target %q", currentPageNumber+1, nextPageTarget) // log the pager lookup
		if nextPageTarget == "" {                                                                                 // when there is no next page link
			log.Printf("no next page link on page %d, so this is the last page", currentPageNumber) // log why we stop
			break                                                                                   // we have reached the end
		} // end of the target check
		pagingFormValues := collectFormFields(currentPageHtml)                                                                                       // start from the page's form data
		pagingFormValues.Set("__EVENTTARGET", nextPageTarget)                                                                                        // tell ASP.NET which pager link was clicked
		pagingFormValues.Set("__EVENTARGUMENT", "")                                                                                                  // the pager sends no extra argument
		pagingResponse, pagingError := crawler.postFormToListingPage(pagingFormValues, "go to the page after page "+strconv.Itoa(currentPageNumber)) // request the next page
		if pagingError != nil {                                                                                                                      // when the request failed
			log.Fatalf("could not load the next page: %v", pagingError) // stop with a clear message
		} // end of the paging error check
		nextPageBytes, _ := io.ReadAll(pagingResponse.Body)                                                                                    // read the new page
		pagingResponse.Body.Close()                                                                                                            // close the body
		currentPageHtml = string(nextPageBytes)                                                                                                // use the new page for the next loop
		log.Printf("received the next page: %d bytes, the pager now shows page %d", len(nextPageBytes), findActivePageNumber(currentPageHtml)) // log what arrived
		if findActivePageNumber(currentPageHtml) <= currentPageNumber {                                                                        // when the page number did not increase
			log.Printf("the page number did not increase, so the last page was already reached") // log why we stop
			break                                                                                // the last page was already reached
		} // end of the progress check
		log.Printf("pausing %s before the next page", crawler.pauseBetweenRequests) // log the pause
		time.Sleep(crawler.pauseBetweenRequests)                                    // be polite to the server
	} // end of the page loop
	log.Printf("finished: %d new PDFs saved, %d skipped, %d failed, folder %s", newFilesSavedCount, skippedCount, failedCount, outputDirectoryName) // final summary
	crawler.logVisitedAddresses()                                                                                                                   // list every address that was visited
	log.Printf("program finished")                                                                                                                  // last log line
} // end of main
