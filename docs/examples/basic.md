# Basic Example

A minimal presentation to get started.

## hello-world.json

```json
{
  "version": "1.0",
  "metadata": {
    "title": "Hello World",
    "author": "Your Name",
    "date": "2025-01-15"
  },
  "theme": {
    "name": "default",
    "aspectRatio": "16:9"
  },
  "slides": [
    {
      "id": "title",
      "type": "title",
      "title": "Hello, PresentationSpec!",
      "widgets": [
        {
          "id": "subtitle",
          "type": "content_block",
          "content": "A simple presentation"
        }
      ]
    },
    {
      "id": "content",
      "type": "content",
      "title": "Key Points",
      "widgets": [
        {
          "id": "points",
          "type": "content_block",
          "content": "## What is PresentationSpec?\n\n- JSON-based presentation format\n- Go-first design with generated schema\n- Multiple output formats (HTML, PDF, PPTX)",
          "format": "markdown"
        }
      ]
    },
    {
      "id": "summary",
      "type": "summary",
      "title": "Thank You",
      "keyMessage": "Get started at github.com/grokify/presentation-spec"
    }
  ]
}
```

## Rendering

Generate HTML:

```bash
presspec render hello-world.json --output ./output
```

Generate PDF:

```bash
presspec render hello-world.json --output ./output --format pdf
```

## File Structure

After rendering:

```
output/
├── index.html
└── assets/
    └── styles.css
```

## Adding Slides

### Content Slide

```json
{
  "id": "features",
  "type": "content",
  "title": "Features",
  "widgets": [
    {
      "id": "feature-list",
      "type": "content_block",
      "content": "- Feature 1\n- Feature 2\n- Feature 3",
      "format": "markdown"
    }
  ]
}
```

### Section Divider

```json
{
  "id": "section-demo",
  "type": "section",
  "title": "Demo Section"
}
```

### Image Slide

```json
{
  "id": "screenshot",
  "type": "content",
  "title": "Product Screenshot",
  "widgets": [
    {
      "id": "hero-image",
      "type": "image",
      "src": "./images/screenshot.png",
      "alt": "Product screenshot",
      "caption": "New dashboard design"
    }
  ]
}
```

## Adding Theme

Customize colors:

```json
{
  "theme": {
    "name": "custom",
    "colors": {
      "primary": "#1a73e8",
      "background": "#ffffff",
      "text": "#1e293b"
    }
  }
}
```

## Adding Speaker Notes

```json
{
  "id": "title",
  "type": "title",
  "title": "Hello, PresentationSpec!",
  "speakerNotes": "Welcome the audience. Introduce yourself."
}
```

Generate with speaker view:

```bash
presspec render hello-world.json --output ./output --speaker-notes
```
