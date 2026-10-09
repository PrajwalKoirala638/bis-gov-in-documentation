<div align="center">

# 📚 BIS Standards Archive

### A PDF backup of the standards published on [www.bis.gov.in](https://www.bis.gov.in/)

![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)
![Format: PDF](https://img.shields.io/badge/Format-PDF-red)
![Source: BIS](https://img.shields.io/badge/Source-bis.gov.in-blue)
![Purpose: Backup](https://img.shields.io/badge/Purpose-Backup%20and%20Archive-orange)
![AI Ready](https://img.shields.io/badge/AI-Pipeline%20Included-purple)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen)

**🏠 Better homes · 💼 Stronger businesses · 🤝 Safer communities · 🤖 AI-ready knowledge**

</div>

---

## 📑 Table of Contents

- [📌 Overview](#-overview)
- [🗂️ Repository Structure](#️-repository-structure)
- [🏛️ About BIS](#️-about-bis)
- [🚀 Quick Start](#-quick-start)
- [👥 Who Is This For?](#-who-is-this-for)
- [🧭 How to Apply a Standard (Step by Step)](#-how-to-apply-a-standard-step-by-step)
- [🏠 Improving Your Home](#-improving-your-home)
- [💼 Improving Your Business](#-improving-your-business)
- [🤝 Improving Your Community](#-improving-your-community)
- [📖 How to Read a Standard](#-how-to-read-a-standard)
- [✅ BIS Marks and Certification](#-bis-marks-and-certification)
- [🤖 AI and Machine Learning Readiness](#-ai-and-machine-learning-readiness)
- [🗺️ Roadmap](#️-roadmap)
- [❓ FAQ](#-faq)
- [⚠️ Disclaimer and Copyright](#️-disclaimer-and-copyright)
- [🛠️ Contributing](#️-contributing)
- [📄 License](#-license)

---

## 📌 Overview

This repository is a **backup archive** of the standards available on the Bureau of Indian Standards (BIS) website, saved in **PDF format**.

> 📁 **All PDF files are stored in the [`PDFs/`](./PDFs) folder.**

**Why does this exist?**

| 🎯 Goal                  | 💡 What it means                                                                                                                                                |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Preservation**         | Websites change, links break, and files get moved or removed. A backup keeps the knowledge available.                                                           |
| **Offline access**       | Download once and read anywhere, including places with poor or no internet.                                                                                     |
| **One searchable place** | Every standard sits in a single folder that is easy to browse, search, and clone.                                                                               |
| **Open learning**        | Students, builders, and small business owners can study the same technical references professionals use.                                                        |
| **AI-ready knowledge**   | The archive is designed to be converted into clean, structured data for search tools and AI assistants (see [AI section](#-ai-and-machine-learning-readiness)). |

---

## 🗂️ Repository Structure

```text
bis-gov-in-documentation/
├── 📁 PDFs/            # All standards in PDF form (the main content)
├── 📄 README.md        # You are here
├── 📄 LICENSE          # MIT license (applies to repository files only)
└── 📄 .gitignore
```

**Suggested additions** (described in the [AI section](#-ai-and-machine-learning-readiness)):

```text
├── 📁 scripts/         # Helper scripts (manifest builder, text extractor)
├── 📁 dataset/         # Generated JSON / JSONL files for search and AI use
└── 📄 manifest.csv     # Index of every PDF: file name, pages, size, checksum
```

**Recommended file naming** for anything you add to `PDFs/`:

```text
IS-<number>-<year>-<Short-Title>.pdf
Example: IS-456-2000-Plain-and-Reinforced-Concrete.pdf
```

---

## 🏛️ About BIS

The **Bureau of Indian Standards (BIS)** is India's national standards body. It develops Indian Standards (IS), and runs product certification and hallmarking schemes.

| 📊 Fact                | Detail                                                                                       |
| ---------------------- | -------------------------------------------------------------------------------------------- |
| **Headquarters**       | New Delhi, India                                                                             |
| **Governing ministry** | Ministry of Consumer Affairs, Food and Public Distribution                                   |
| **Legal basis**        | BIS Act, 2016 (which replaced the earlier BIS Act, 1986)                                     |
| **Predecessor**        | Indian Standards Institution (ISI), founded in 1947                                          |
| **What it publishes**  | Specifications, test methods, codes of practice, guides, and terminology                     |
| **Scale**              | Well over 20,000 Indian Standards across most sectors of the economy                         |
| **International link** | Represents India in ISO and IEC, and adopts many international standards as IS/ISO or IS/IEC |
| **Official website**   | [www.bis.gov.in](https://www.bis.gov.in/)                                                    |

**Types of documents you will encounter:**

| Type                         | Purpose                                                                        |
| ---------------------------- | ------------------------------------------------------------------------------ |
| 📏 **Product specification** | Defines what a product must be (materials, dimensions, performance, marking).  |
| 🧪 **Test method**           | Explains exactly how to measure or verify a property.                          |
| 🏗️ **Code of practice**      | Describes recommended good practice for design, construction, or installation. |
| 📘 **Guide / handbook**      | Gives background, explanations, and worked guidance.                           |
| 🔤 **Terminology / symbols** | Standardises vocabulary so everyone means the same thing.                      |

---

## 🚀 Quick Start

**1️⃣ Browse and download one file**

Open the [`PDFs/`](./PDFs) folder, click a file, then choose **Download raw file**.

**2️⃣ Clone the whole archive**

```bash
git clone https://github.com/PrajwalKoirala638/bis-gov-in-documentation.git
cd bis-gov-in-documentation/PDFs
```

**3️⃣ Download as a ZIP (no Git needed)**

Click the green **Code** button, then **Download ZIP**.

**4️⃣ Search inside all PDFs (command line)**

```bash
# Install pdfgrep, then search every PDF for a keyword
pdfgrep -ril "earthquake" PDFs/          # list files that mention it
pdfgrep -rin "minimum cover" PDFs/       # show matching lines with line numbers
```

**5️⃣ Search by file name**

Use the **Go to file** button on GitHub (or press `t` on the repository page) and type a standard number such as `456`.

---

## 👥 Who Is This For?

| 👤 Audience                            | 🔧 How they can use these standards                                         |
| -------------------------------------- | --------------------------------------------------------------------------- |
| 🏡 **Homeowners**                      | Check that materials and construction methods meet recognised requirements. |
| 👷 **Builders and masons**             | Follow proven practice for concrete, masonry, and finishing.                |
| 📐 **Architects and engineers**        | Reference design, loading, seismic, and material requirements.              |
| ⚡ **Electricians**                    | Follow wiring, earthing, and electrical safety practice.                    |
| 🚰 **Plumbers**                        | Choose suitable pipes and fittings and install them correctly.              |
| 🏭 **Manufacturers**                   | Design products to a recognised specification and test them properly.       |
| 🛒 **Buyers and procurement teams**    | Write clear specifications and verify supplier claims.                      |
| 🚀 **Entrepreneurs and startups**      | Understand quality and compliance requirements before launching a product.  |
| 🎓 **Students and teachers**           | Learn how technical standards are written and applied in practice.          |
| 🧑‍🤝‍🧑 **Community leaders and NGOs**      | Plan safer schools, clinics, water systems, and shared buildings.           |
| 🔍 **Journalists and watchdog groups** | Compare public works and product claims with the published requirements.    |
| 🤖 **Developers and researchers**      | Build search tools, assistants, and datasets from the documents.            |

---

## 🧭 How to Apply a Standard (Step by Step)

Anyone can turn a standard into practical improvement by following this simple workflow:

| Step  | Action                                  | Example                                                                              |
| :---: | --------------------------------------- | ------------------------------------------------------------------------------------ |
| **1** | 🎯 **Define your need**                 | "I am building a two-storey house." / "I want to sell water pipes."                  |
| **2** | 🔎 **Find the relevant standard**       | Search the `PDFs/` folder or use the keyword table in this README.                   |
| **3** | 📅 **Check the edition**                | Confirm the document is the latest version on [bis.gov.in](https://www.bis.gov.in/). |
| **4** | 📖 **Read the Scope first**             | Make sure the standard actually covers your situation.                               |
| **5** | 📝 **List the requirements**            | Turn "shall" statements into a checklist for your project.                           |
| **6** | 🧪 **Verify with tests or documents**   | Ask for test reports, certificates, or inspection records.                           |
| **7** | 📂 **Keep records**                     | Save photos, receipts, and test reports as proof of compliance.                      |
| **8** | 👨‍🔧 **Involve a qualified professional** | Use a licensed engineer or inspector for anything safety-critical.                   |

> 💡 **Tip:** The word **"shall"** in a standard marks a requirement. **"Should"** marks a recommendation.

---

## 🏠 Improving Your Home

Standards turn "good enough" into "built properly." Use the table below as a starting map.

| 🧱 Area                      | ✅ What to check                                  | 📚 Example standards to look for                    |
| ---------------------------- | ------------------------------------------------- | --------------------------------------------------- |
| **Structure and foundation** | Concrete quality, reinforcement, load assumptions | IS 456, IS 875 (loads), IS 1786 (steel bars)        |
| **Earthquake safety**        | Seismic design and ductile detailing              | IS 1893, IS 13920, IS 4326                          |
| **Cement and aggregates**    | Correct cement grade and clean, graded aggregate  | IS 269, IS 8112, IS 12269, IS 1489, IS 383          |
| **Walls and masonry**        | Brick and block strength and size                 | IS 1077, IS 2185                                    |
| **Electrical wiring**        | Cable type, sockets, circuit protection           | IS 732, IS 694, IS 1293                             |
| **Earthing**                 | Safe earthing to prevent shocks and fires         | IS 3043                                             |
| **Water and plumbing**       | Suitable pipes, safe drinking water               | IS 4985, IS 1239, IS 10500                          |
| **Water testing**            | How to sample and test water                      | IS 3025 (series)                                    |
| **Fire safety**              | Extinguishers and exit planning                   | IS 2190, IS 15683, NBC Part 4                       |
| **Lighting and appliances**  | Safe, efficient products                          | IS 16102 (LED lamps), IS 302 (household appliances) |
| **Soil and site**            | Soil testing before building                      | IS 2720 (series)                                    |

> 🔖 _Standard numbers above are examples to help you search. Years and editions change, so always use the exact files in the `PDFs/` folder and confirm the latest edition on the official site._

### 🏡 Home Quick-Win Checklist

- [ ] ✅ Look for the **BIS Standard Mark** on cement, steel bars, cables, switches, and pipes
- [ ] 🧾 Ask suppliers for **test certificates** and keep the invoices
- [ ] 📐 Share the relevant standard with your **mason or contractor** so expectations are clear
- [ ] ⚡ Make sure the home has proper **earthing** and appropriately rated protection devices
- [ ] 🔥 Install and maintain a suitable **fire extinguisher** and plan an escape route
- [ ] 💧 **Test your drinking water** and compare results with the drinking water specification
- [ ] 🏗️ Follow **earthquake-resistant** detailing if you live in a seismic zone
- [ ] 📸 **Photograph key stages** (reinforcement before concrete, wiring before plastering)

---

## 💼 Improving Your Business

Standards give businesses a shared language of quality. Used well, they reduce waste, build trust, and open doors.

### 🌟 Business Benefits

| 💎 Benefit                    | 📈 How standards help                                                              |
| ----------------------------- | ---------------------------------------------------------------------------------- |
| **Higher product quality**    | Design to a recognised specification and verify output with standard test methods. |
| **Customer trust**            | A named standard is a verifiable promise, not just a marketing claim.              |
| **Fewer defects and returns** | Tolerances and test methods catch problems before products ship.                   |
| **Cleaner procurement**       | Reference standards in purchase orders and contracts to avoid disputes.            |
| **Safer workplaces**          | Electrical, fire, and equipment standards reduce accidents and downtime.           |
| **Smoother compliance**       | Knowing requirements early makes audits and certification easier.                  |
| **Better training**           | Standards are strong reference material for onboarding staff.                      |
| **Market access**             | Buyers, tenders, and regulators often expect conformity to recognised standards.   |

### 🏭 By Business Type

| 🏢 Sector                        | 🔧 How to use the standards                                                         |
| -------------------------------- | ----------------------------------------------------------------------------------- |
| **Construction and real estate** | Specify materials, structural design, and fire safety in contracts and inspections. |
| **Manufacturing**                | Build in-house quality checklists from product specs and test methods.              |
| **Retail and distribution**      | Verify supplier claims and spot unmarked or substandard goods.                      |
| **Electrical and electronics**   | Check product safety requirements before importing or selling.                      |
| **Hospitality and food service** | Use fire safety, water quality, and building guidance for premises.                 |
| **Startups**                     | Learn compliance expectations early, before expensive redesigns.                    |
| **Training and consulting**      | Create courses, audits, and checklists based on the documents.                      |

### 🪜 Business Action Plan

1. 🔍 **Identify** which standards apply to your product or premises.
2. 📋 **Gap check:** compare your current practice with the requirements.
3. 🛠️ **Fix** the gaps, starting with safety-critical items.
4. 🧪 **Test** products using the standard test methods or an accredited lab.
5. 🏷️ **Certify** where required, following the official BIS process.
6. 🔄 **Review** regularly, because standards are revised over time.

> 🌐 Many businesses also adopt management system standards such as **IS/ISO 9001** (quality), **IS/ISO 14001** (environment), and **IS/ISO 45001** (occupational health and safety). Check whether the relevant ones are in the `PDFs/` folder.

---

## 🤝 Improving Your Community

Shared infrastructure affects everyone. Standards help communities plan, build, and hold projects to a clear benchmark.

| 🌍 Community need                  | 🛠️ How the standards help                                                                        |
| ---------------------------------- | ------------------------------------------------------------------------------------------------ |
| 🏫 **Schools, clinics, and halls** | Plan structurally safe, fire-safe, and accessible public buildings.                              |
| 💧 **Clean water**                 | Use drinking water specifications and test methods to design and monitor shared supplies.        |
| 🚰 **Sanitation**                  | Follow plumbing and drainage guidance for safer, healthier neighbourhoods.                       |
| 🌪️ **Disaster preparedness**       | Apply hazard-resistant construction guidance to reduce harm from earthquakes, storms, and fires. |
| 🎓 **Skills and livelihoods**      | Train masons, electricians, plumbers, and technicians using trusted references.                  |
| 🔍 **Transparency**                | Allow residents and committees to compare public works with published requirements.              |
| 🛒 **Consumer protection**         | Teach people how to spot genuine, certified products.                                            |
| 📚 **Education**                   | Give teachers and students real technical documents for engineering and vocational courses.      |

### 🌱 Community Project Ideas

- 🧑‍🏫 Run a **workshop** for local masons on concrete and masonry basics
- 💦 Start a **water-quality testing** drive for wells and taps
- 🔌 Host an **electrical safety** awareness day
- 🏚️ Create a **school safety checklist** for parents and teachers
- 📱 Build a **local resource library** (offline copies shared on a USB drive or local server)
- 🗣️ Translate key sections into **local languages** for wider reach (with attribution and care)

---

## 📖 How to Read a Standard

Most Indian Standards follow a similar layout:

|  #  | Section                  | 🔎 What to look for                                           |
| :-: | ------------------------ | ------------------------------------------------------------- |
|  1  | **Foreword**             | Why the standard exists and what changed in this revision.    |
|  2  | **Scope**                | What it covers, and what it does not. Always read this first. |
|  3  | **References**           | Other standards you also need.                                |
|  4  | **Terminology**          | Definitions of key terms.                                     |
|  5  | **Requirements**         | The "shall" statements: materials, dimensions, performance.   |
|  6  | **Sampling and testing** | How many samples to take and how to test them.                |
|  7  | **Marking and packing**  | Labelling and conformity marking.                             |
|  8  | **Annexes**              | Extra methods, tables, and guidance.                          |

---

## ✅ BIS Marks and Certification

BIS runs several schemes that help consumers recognise quality. Check the official site for the current list of products covered and the latest rules.

| 🏷️ Scheme                                | 📝 What it is                                                                          |
| ---------------------------------------- | -------------------------------------------------------------------------------------- |
| **BIS Standard Mark (ISI mark)**         | Shows a product conforms to the relevant Indian Standard under a BIS licence.          |
| **Hallmarking**                          | Purity marking for gold and silver jewellery and artefacts.                            |
| **Compulsory Registration Scheme (CRS)** | Registration requirement for certain electronic and IT products.                       |
| **BIS Care app**                         | Official app for checking licence and registration details and for reporting concerns. |

---

## 🤖 AI and Machine Learning Readiness

> ⚠️ **Honest status:** Right now the archive consists of **raw PDF files**. PDFs are designed for people to read, not for machines to learn from. They contain tables, multi-column layouts, scanned pages, and formulas that need cleaning first. The sections below explain **how to convert these PDFs into a structured, AI-ready dataset** and what you can build with it.

### 🧠 What Can You Build?

| 🤖 Application                                    | 📝 Description                                                                              |
| ------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| 💬 **Standards Q&A assistant**                    | A chatbot that answers questions and **cites the exact standard and page**.                 |
| 🔎 **Semantic search**                            | Find the right clause by meaning, not just by exact keywords.                               |
| ✅ **Compliance checker**                         | Describe a project and get a list of potentially relevant standards and checklists.         |
| 📊 **Requirement extraction**                     | Pull limits, tolerances, and tables into structured data (CSV or JSON).                     |
| 🌐 **Plain-language and multilingual explainers** | Summarise technical clauses for non-experts and translate with human review.                |
| 🎓 **Learning tools**                             | Generate quizzes, flashcards, and study guides from the standards.                          |
| 🧪 **Evaluation datasets**                        | Create question-and-answer benchmarks to test how well AI tools handle technical documents. |

### 🔄 Recommended Data Pipeline

```text
 PDFs/  ──►  Manifest  ──►  Text extraction  ──►  Cleaning  ──►  Chunking  ──►  Embeddings  ──►  Search / AI app
 (raw)     (index +        (page by page,       (headers,      (sections,     (vector DB)       (with citations)
           checksums)       OCR if scanned)      footers)       with metadata)
```

| 🪜 Stage               | 🔧 Details                                                             | 🧰 Useful tools       |
| ---------------------- | ---------------------------------------------------------------------- | --------------------- |
| **1. Manifest**        | Record file name, pages, size, and checksum for every PDF.             | Python, `hashlib`     |
| **2. Extraction**      | Pull text page by page and keep page numbers.                          | PyMuPDF, pdfplumber   |
| **3. OCR (if needed)** | Scanned pages need optical character recognition.                      | OCRmyPDF, Tesseract   |
| **4. Cleaning**        | Remove repeated headers, footers, and page artefacts.                  | Python, regex         |
| **5. Chunking**        | Split by section or clause, not at random character counts.            | Custom rules          |
| **6. Metadata**        | Attach standard number, title, year, section, and page to every chunk. | JSON / JSONL          |
| **7. Embeddings**      | Convert chunks into vectors for semantic search.                       | sentence-transformers |
| **8. Vector store**    | Store and query vectors.                                               | FAISS, Chroma         |
| **9. Answering**       | Retrieve relevant chunks and generate an answer **with citations**.    | Any LLM               |

### 🧾 Suggested Data Schema

**Document-level record** (`dataset/documents.jsonl`, one line per PDF):

```json
{
  "doc_id": "IS-456-2000",
  "title": "Plain and Reinforced Concrete - Code of Practice",
  "standard_number": "IS 456",
  "year": 2000,
  "category": "Civil Engineering",
  "language": "en",
  "source_file": "PDFs/IS-456-2000-Plain-and-Reinforced-Concrete.pdf",
  "source_url": "https://www.bis.gov.in/",
  "pages": 114,
  "sha256": "<checksum>",
  "verified_current_edition": false
}
```

**Chunk-level record** (`dataset/chunks.jsonl`, one line per section or passage):

```json
{
  "chunk_id": "IS-456-2000_p12_c03",
  "doc_id": "IS-456-2000",
  "page_start": 12,
  "page_end": 12,
  "section": "Scope",
  "text": "<cleaned text of the passage>",
  "token_count": 312
}
```

> 🔖 The values above are **illustrative examples** of the format, not statements about any specific file.

### 🧰 Starter Scripts

<details>
<summary><b>📋 1. Build a manifest of all PDFs (click to expand)</b></summary>

```python
# scripts/build_manifest.py
# pip install pymupdf
import csv
import hashlib
import pathlib

import fitz  # PyMuPDF

rows = []
for pdf in sorted(pathlib.Path("PDFs").glob("*.pdf")):
    data = pdf.read_bytes()
    with fitz.open(pdf) as doc:
        pages = doc.page_count
    rows.append({
        "file": pdf.name,
        "pages": pages,
        "size_bytes": len(data),
        "sha256": hashlib.sha256(data).hexdigest(),
    })

if rows:
    with open("manifest.csv", "w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=rows[0].keys())
        writer.writeheader()
        writer.writerows(rows)
    print(f"Wrote manifest.csv with {len(rows)} files")
else:
    print("No PDFs found in PDFs/")
```

</details>

<details>
<summary><b>📄 2. Extract text page by page into JSONL (click to expand)</b></summary>

```python
# scripts/extract_text.py
# pip install pymupdf
import json
import pathlib

import fitz  # PyMuPDF

pathlib.Path("dataset").mkdir(exist_ok=True)

with open("dataset/pages.jsonl", "w", encoding="utf-8") as out:
    for pdf in sorted(pathlib.Path("PDFs").glob("*.pdf")):
        with fitz.open(pdf) as doc:
            for page_number, page in enumerate(doc, start=1):
                text = page.get_text("text").strip()
                if text:  # pages with no text may be scans and need OCR
                    record = {"doc_id": pdf.stem, "page": page_number, "text": text}
                    out.write(json.dumps(record, ensure_ascii=False) + "\n")

print("Done. See dataset/pages.jsonl")
```

</details>

<details>
<summary><b>🔍 3. OCR scanned PDFs (click to expand)</b></summary>

```bash
# Install OCRmyPDF, then add a searchable text layer to a scanned PDF
ocrmypdf --skip-text input.pdf output.pdf
```

</details>

### ✅ Best Practices for AI Use

- 📌 **Always cite the source.** Every answer should point to the standard number, edition, and page.
- 🧑‍⚖️ **Keep a human in the loop.** AI can misread tables, units, and clauses. Verify before acting.
- 🚫 **Never rely on AI alone for safety-critical decisions** such as structural, electrical, or fire safety design.
- 🔢 **Version everything.** Record the edition and the date you extracted the text.
- 🧪 **Evaluate.** Build a small set of questions with known answers and test your system regularly.
- 🔄 **Re-check against the official site.** Standards are revised, and an AI tool is only as current as its data.
- 📐 **Handle tables and formulas carefully.** Plain text extraction often breaks them, so review the output.

### ⚖️ Legal and Ethical Note on AI Use

BIS standards are copyrighted publications. **Using them to train or fine-tune a model, or redistributing extracted text, may require permission from BIS** or the relevant rights holder. Building a private search or question-answering tool for your own use is a different situation from publishing a dataset or training a public model. Check BIS's terms of use, and consult a legal professional if you plan to share or commercialise derived data or models.

---

## 🗺️ Roadmap

- [x] 📁 Upload standards as PDFs to `PDFs/`
- [x] 📝 Write documentation
- [ ] 📋 Generate `manifest.csv` (file, pages, size, checksum)
- [ ] 🔤 Standardise file names (`IS-<number>-<year>-<Title>.pdf`)
- [ ] 🏷️ Add category tags (civil, electrical, water, fire, and so on)
- [ ] 📄 Extract text into `dataset/pages.jsonl`
- [ ] 🔍 Add OCR for scanned documents
- [ ] 🧩 Create `chunks.jsonl` with section-level metadata
- [ ] 🔎 Build a searchable index
- [ ] 💬 Prototype a citation-first Q&A assistant
- [ ] 🌐 Add plain-language summaries for popular standards
- [ ] ✅ Mark each document with the date its edition was last verified

---

## ❓ FAQ

<details>
<summary><b>Is this an official BIS repository?</b></summary>

No. This is an unofficial personal backup. It is not affiliated with, endorsed by, or maintained by BIS or the Government of India.

</details>

<details>
<summary><b>Are these documents up to date?</b></summary>

Not guaranteed. Standards are revised and withdrawn over time. Always confirm the current edition at [www.bis.gov.in](https://www.bis.gov.in/) before relying on a document.

</details>

<details>
<summary><b>Can I use these for a real construction or business project?</b></summary>

They are excellent reference material, but for structural, electrical, or other safety-critical work you should involve a qualified, licensed professional and use the latest official edition.

</details>

<details>
<summary><b>Can I build an AI tool with these files?</b></summary>

Technically yes, and this README shows a recommended pipeline. Legally, check BIS's terms and copyright position first, especially before training a model or publishing derived data. See the [legal note](#️-legal-and-ethical-note-on-ai-use).

</details>

<details>
<summary><b>How do I find the right standard?</b></summary>

Use GitHub's **Go to file** search with the standard number, search with `pdfgrep` for keywords, or use the topic tables in this README to find likely standard numbers first.

</details>

<details>
<summary><b>How can I help?</b></summary>

Open an issue for missing or outdated files, improve the documentation, or submit a pull request with new PDFs, tags, or scripts.

</details>

---

## ⚠️ Disclaimer and Copyright

- 🗄️ **Backup only.** This repository is an unofficial archive created for preservation and learning.
- 🏛️ **Not affiliated with BIS.** It is not endorsed by or connected to the Bureau of Indian Standards or the Government of India.
- 📅 **Check the latest version.** Documents may be superseded, amended, or withdrawn. Verify at [www.bis.gov.in](https://www.bis.gov.in/).
- 🧑‍🔧 **Not professional advice.** Content is reference material only. Consult qualified professionals for engineering, electrical, structural, legal, or safety decisions.
- ©️ **Copyright.** The PDF content belongs to **BIS** and/or its respective rights holders. The MIT license in this repository applies only to the repository's own original files (such as this README and any scripts), **not** to the standards themselves.
- 📬 **Takedown requests.** If you are a rights holder and want something removed, please open an issue and it will be addressed promptly.
- 🛒 **Official copies.** For authoritative copies, purchases, and certification, always go to BIS directly.

---

## 🛠️ Contributing

Contributions are welcome. 🙌

1. 🍴 **Fork** the repository.
2. 🌿 **Create a branch:** `git checkout -b feature/my-improvement`
3. ✏️ **Make your changes** (add PDFs to `PDFs/`, improve docs, add scripts).
4. 💾 **Commit:** `git commit -m "Add: short description"`
5. 🚀 **Push and open a Pull Request.**

**Guidelines:**

- Use the file naming format `IS-<number>-<year>-<Short-Title>.pdf`
- Only add documents that are publicly available from the official source
- Do not commit personal data or unrelated files
- Keep documentation clear, accurate, and professional

---

## 📄 License

The repository's own original files are released under the [MIT License](./LICENSE).
The PDF documents remain the property of their original publisher (see the [disclaimer](#️-disclaimer-and-copyright)).

---

<div align="center">

### 🔗 Official Source

**Bureau of Indian Standards:** [https://www.bis.gov.in/](https://www.bis.gov.in/)

⭐ If this archive helps you, consider starring the repository and sharing it with your community. ⭐

</div>
