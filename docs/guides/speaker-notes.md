# Speaker Notes

This guide covers using speaker notes in presentations.

## Overview

Speaker notes provide presenter-only content that doesn't appear on slides. PresentationSpec supports:

- Per-slide speaker notes
- Dedicated speaker view with timer
- Keyboard navigation

## Adding Speaker Notes

Add notes to any slide:

```json
{
  "id": "overview",
  "type": "content",
  "title": "Product Overview",
  "speakerNotes": "Start with the problem statement. Mention customer feedback from Q4 survey.",
  "widgets": [...]
}
```

### Markdown Support

Speaker notes support markdown formatting:

```json
{
  "speakerNotes": "## Key Points\n\n- Revenue up 15%\n- User growth exceeds target\n- **Emphasize** the Q2 launch date"
}
```

## Generating Speaker View

Generate a separate speaker notes HTML file:

```bash
presspec render presentation.json --output ./output --speaker-notes
```

This creates:

```
output/
├── index.html          # Main presentation
├── speaker-notes.html  # Speaker view
└── assets/
```

## Speaker View Features

The speaker view includes:

### Timer

- Elapsed time display
- Start/pause/reset controls

### Slide Preview

- Current slide thumbnail
- Next slide preview

### Notes Panel

- Current slide notes (rendered markdown)
- Large, readable text

### Navigation

| Key | Action |
|-----|--------|
| `→` / `Space` | Next slide |
| `←` | Previous slide |
| `Home` | First slide |
| `End` | Last slide |

## Dual-Screen Setup

For presentations:

1. Open `index.html` on the projector/external display
2. Open `speaker-notes.html` on your laptop screen
3. Navigate from either window (they stay in sync via localStorage)

## Example

```json
{
  "version": "1.0",
  "metadata": {
    "title": "Q1 Review"
  },
  "slides": [
    {
      "id": "title",
      "type": "title",
      "title": "Q1 2025 Review",
      "speakerNotes": "Welcome everyone. This is our quarterly review covering January through March."
    },
    {
      "id": "agenda",
      "type": "content",
      "title": "Agenda",
      "speakerNotes": "## Topics\n\n1. Financial highlights (5 min)\n2. Product updates (10 min)\n3. Q2 roadmap (5 min)\n4. Q&A (10 min)\n\n**Total: 30 minutes**",
      "widgets": [...]
    },
    {
      "id": "revenue",
      "type": "dashboard",
      "title": "Revenue",
      "speakerNotes": "Key message: Revenue up 15% YoY.\n\n- Driven by enterprise segment\n- APAC grew 25%\n- Pause for questions if needed"
    },
    {
      "id": "summary",
      "type": "summary",
      "title": "Summary",
      "speakerNotes": "Recap the three key points:\n\n1. Strong financial performance\n2. Product momentum\n3. Clear Q2 priorities\n\nThank the team for their contributions."
    }
  ]
}
```

## Best Practices

1. **Keep notes concise** - Bullet points over paragraphs
2. **Include timing cues** - Note when to pause or ask questions
3. **Add data reminders** - Key numbers you want to emphasize
4. **Use markdown headers** - Organize longer notes
5. **Practice with speaker view** - Rehearse with the timer running
