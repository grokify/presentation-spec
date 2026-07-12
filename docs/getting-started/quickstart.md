# Quick Start

Create your first presentation in 5 minutes.

## Step 1: Create a Presentation Spec

Create a file called `hello.json`:

```json
{
  "version": "1.0",
  "metadata": {
    "title": "Hello PresentationSpec",
    "description": "My first presentation",
    "author": "Your Name"
  },
  "theme": {
    "name": "corporate",
    "aspectRatio": "16:9"
  },
  "slides": [
    {
      "id": "title",
      "type": "title",
      "title": "Hello PresentationSpec",
      "widgets": [
        {
          "id": "subtitle",
          "type": "content_block",
          "content": "A Go-first presentation library"
        }
      ]
    },
    {
      "id": "features",
      "type": "content",
      "title": "Features",
      "keyMessage": "Everything you need for data-driven presentations",
      "widgets": [
        {
          "id": "feature-list",
          "type": "checklist",
          "items": [
            {"text": "Go-first design", "checked": true},
            {"text": "Multiple output formats", "checked": true},
            {"text": "Rich widget library", "checked": true},
            {"text": "Data integration", "checked": true}
          ]
        }
      ]
    },
    {
      "id": "metrics",
      "type": "dashboard",
      "title": "Sample Dashboard",
      "widgets": [
        {
          "id": "metric1",
          "type": "metric_card",
          "title": "Users",
          "value": "10K",
          "status": "good",
          "trend": "up",
          "trendValue": "+25%"
        },
        {
          "id": "metric2",
          "type": "metric_card",
          "title": "Revenue",
          "value": "1.2M",
          "unit": "$",
          "status": "good"
        }
      ]
    }
  ]
}
```

## Step 2: Render to HTML

```bash
render hello.json -o hello-output/
```

This creates:

```
hello-output/
├── index.html
└── assets/
    ├── style.css
    └── navigation.js
```

## Step 3: View the Presentation

Open `hello-output/index.html` in a browser.

Use arrow keys or click the navigation buttons to move between slides.

## Step 4: Use Watch Mode

For rapid iteration, use watch mode:

```bash
render hello.json -o hello-output/ --watch
```

Edit `hello.json` and the output rebuilds automatically.

## Step 5: Start the Dev Server

For the best development experience, use the dev server:

```bash
render serve hello.json
```

This starts a server at `http://localhost:8080` with live reload.

## Step 6: Export to PDF

```bash
render hello.json --format pdf -o hello.pdf
```

## Step 7: Add Speaker Notes

Add speaker notes and generate the presenter view:

```json
{
  "id": "metrics",
  "type": "dashboard",
  "title": "Sample Dashboard",
  "speakerNotes": "Emphasize the 25% user growth this quarter.",
  "widgets": [...]
}
```

Generate with speaker notes view:

```bash
render hello.json -o hello-output/ --speaker-notes
```

Open `hello-output/speaker.html` for the presenter view with timer and notes.

## Next Steps

- [CLI Reference](cli.md) - All commands and options
- [Widgets Reference](../specification/widgets.md) - All 12 widget types
- [Dashboard Guide](../guides/dashboards.md) - Create executive dashboards
- [Examples](../examples/basic.md) - More example presentations
