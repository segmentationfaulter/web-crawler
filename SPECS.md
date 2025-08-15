web-crawler/SPECS.md
# Web Scraper Project Specifications

## Project Overview
Create a web crawler in Go that:
1. Starts from a specified root URL
2. Discovers all links on the page
3. Follows links to crawl entire websites
4. Respects website politeness policies
5. Outputs crawl results in a structured format

## Features
### Core Functionality
- **URL Discovery**: Parse HTML to extract all `<a href>` links
- **Breadth-First Crawling**: Implement BFS algorithm to traverse links
- **Domain Limitation**: Option to stay within the same domain
- **Politeness Policy**: Respect `robots.txt` and crawl-delay directives
- **Concurrency**: Use goroutines for parallel crawling
- **Redirect Handling**: Properly handle HTTP redirects (3xx status codes) by following the `Location` header and updating the URL accordingly. Limit the number of redirects to avoid loops.

### Output Requirements
- **Sitemap Generation**: Output hierarchical site structure
- **Page Metadata**: Record for each page:
  - URL
  - Title
  - HTTP status code
  - Number of links found
  - Timestamp of crawl
- **Error Handling**: Log failed requests and invalid URLs

## Technical Specifications
### Dependencies
- Go standard libraries (`net/http`, `golang.org/x/net/html`)
- Third-party packages:
  - `github.com/temoto/robotstxt` (robots.txt parsing)
  - `github.com/PuerkitoBio/goquery` (optional for HTML parsing)

### Components
1. **Crawler Engine**
   - BFS queue implementation
   - Visited URL tracking (thread-safe)
   - Rate limiting mechanism

2. **HTTP Fetcher**
   - Custom HTTP client with timeout
   - User-Agent rotation
   - Response caching
   - Redirect following with max depth (default 5)
   - Handling of 301 (Permanent), 302 (Found), 307 (Temporary), and 308 (Permanent) status codes

3. **Parser**
   - HTML link extraction
   - Relative to absolute URL conversion
   - Content type filtering

4. **Robots.txt Handler**
   - Fetch and parse robots.txt
   - Check path allowance before crawling

### Configuration
- Max depth level
- Max pages to crawl
- Request delay (ms)
- User agent string
- Output format (JSON/text)

## Setup & Execution
```sh
go run main.go -url https://example.com -depth 3 -output sitemap.json
```

## Milestones
1. Basic single-page crawler
2. Recursive BFS implementation
3. Robots.txt integration
4. Concurrent crawling with worker pool
5. Output formatting
6. Error handling and logging

## Constraints
- Avoid crawling non-HTML content
- Respect website's `robots.txt` rules
- Implement politeness delay (default 200ms)
- Handle at least 10 common HTTP status codes