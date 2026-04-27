# test-site

This repository is a working site built with `press`.

It serves two purposes:

- a development sandbox for building and refining site structure
- a reference implementation of how to integrate `press` in a site repository

---

## Overview

The site is built from:

- Markdown content in `content/`
- templ components in `templates/`
- renderer adapters in `internal/site/`

The build process is handled by `press`, which is responsible for:

- discovering and parsing content (posts and static pages)
- generating routes
- rendering templates
- writing output and syncing assets

---

## Development

### Build

```bash
just build-site
```

Builds the static site output in `build/public/`.

```bash
just build-all
```

Builds the site and any associated binaries (e.g. CLI tools) for the repository.

### Build and serve

```bash
just serve
```

Builds the site and starts a local HTTP server.

### Live development

```bash
just serve-live
```

Uses `air` to rebuild and restart the server on file changes.

---

## Structure

```
content/
  posts/
    my-post/
      index.md
      media/
  pages/
    about/
      index.md

templates/
  base.templ
  home.templ
  blog_index.templ
  blog_post.templ

internal/site/
  renderers.go
```

- **content/**: Markdown content (posts and static pages)
- **templates/**: templ components for HTML rendering
- **internal/site/**: adapters connecting templates to `press`

---

## Rendering

The site provides renderers to `press`:

```go
func Renderers() press.Renderers
```

These map structured page data to templ components.

---

## Content

### Posts

Each post lives in:

```
content/posts/{slug}/index.md
```

With frontmatter:

```yaml
---
title: My Post
slug: my-post
date: 2026-04-25
---
```

Optional media can be placed in:

```
content/posts/{slug}/media/
```

### Static pages

Each static page lives in:

```
content/pages/{slug}/index.md
```

With frontmatter:

```yaml
---
title: My Page
slug: my-page
---
```



---

## Output

The generated site is written to:

```
build/public/
```

With structure:

```
/
  index.html
  blog/
    index.html
    {slug}/
      index.html
      media/
  {slug}/
    index.html
  assets/
```

---

## Deployment

This repository is currently deployed via GitHub Pages as a project site.

This means URLs are served under:

```
/test-site/
```

Some root-relative links may not function correctly in this environment.

Local development (`just serve`) reflects the intended deployment structure.

---

## Adapting This Repository

This repository is intended to be cloned and adapted for new sites (e.g. `spc-dev`, `spc-music`).

After cloning, update the following:

### 1. Site metadata

Update the `SiteData` passed to `press.Build`:

```go
siteData := press.SiteData{
	Title:         "Your Site Name",
	StylesheetURL: "/assets/css/styles.css",
}
```

### 2. Navigation

Update `templates/nav.templ` to reflect the structure of the site:

```html
<nav>
    <a href="/">Home</a>
    <a href="/blog/">Blog</a>
</nav>
```

Static pages (e.g. `/about/`, `/teaching/`) must be linked manually. Add or remove sections as needed.

### 3. Content

Replace or remove the sample content:

```
content/posts/test-post/
```

Add new posts under:

```
content/posts/{slug}/index.md
```

### 4. Templates

Update templates in `templates/`:

- `base.templ`: global layout (head, metadata, styles)
- `home.templ`: homepage content
- `blog_index.templ`: list view of posts
- `blog_post.templ`: individual post layout

### 5. Domain/Deployment

For custom domains (e.g. `spc-dev.com`, `spc-music.com`):

- No changes to routes are required
- Root-relative URLs (`/blog/`, `/assets/...`) will work as expected

If using GitHub Pages project URLs (`/repo-name/`), some links may need adjustment or can be ignored during development (use `just serve` for preview instead).

### 6. Project naming

Update:

- repo name
- module path in go.mod
- any reference to test-site in scripts or docs

### 7. Styling

Replace or extend `static/css/styles.css`

### Notes

`press` defines the content model and build pipeline; this repository defines presentation and content. Preserve this boundary when adapting.

