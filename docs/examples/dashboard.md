# Dashboard Example

A full dashboard presentation with KPIs, charts, and data.

## launch-readiness.json

```json
{
  "version": "1.0",
  "metadata": {
    "title": "Launch Readiness Review",
    "description": "Q1 2025 product launch readiness",
    "author": "Product Team",
    "date": "2025-01-15",
    "tags": ["launch", "readiness", "q1-2025"]
  },
  "theme": {
    "name": "corporate",
    "density": "normal",
    "aspectRatio": "16:9",
    "brand": {
      "logo": "/assets/logo.png",
      "logoPosition": "top-right",
      "companyName": "Acme Corp"
    },
    "colors": {
      "primary": "#1a73e8",
      "secondary": "#34a853",
      "background": "#ffffff",
      "surface": "#f8f9fa"
    }
  },
  "narrative": {
    "storyline": "Launch readiness assessment showing strong progress with manageable risks",
    "sections": [
      {"id": "status", "title": "Current Status"},
      {"id": "risks", "title": "Risks & Blockers"},
      {"id": "next-steps", "title": "Next Steps"}
    ]
  },
  "slides": [
    {
      "id": "title",
      "type": "title",
      "title": "Launch Readiness Review",
      "widgets": [
        {
          "id": "subtitle",
          "type": "content_block",
          "content": "Q1 2025 | Product Team"
        }
      ],
      "speakerNotes": "Welcome everyone to our launch readiness review."
    },
    {
      "id": "dashboard",
      "sectionId": "status",
      "type": "dashboard",
      "title": "Launch Metrics",
      "keyMessage": "On track for launch with 85% readiness",
      "layout": {
        "template": "kpi_row_content_grid",
        "regions": [
          {"id": "kpi-row", "role": "kpi"},
          {"id": "main", "role": "content", "columns": 2},
          {"id": "sidebar", "role": "sidebar", "columns": 1}
        ]
      },
      "widgets": [
        {
          "id": "readiness",
          "type": "metric_card",
          "regionId": "kpi-row",
          "title": "Overall Readiness",
          "value": "85%",
          "status": "good",
          "trend": "up",
          "trendValue": "+5%",
          "target": "80%"
        },
        {
          "id": "features",
          "type": "metric_card",
          "regionId": "kpi-row",
          "title": "Features Complete",
          "value": "12/14",
          "status": "good",
          "trend": "up",
          "trendValue": "+2"
        },
        {
          "id": "bugs",
          "type": "metric_card",
          "regionId": "kpi-row",
          "title": "Open Bugs",
          "value": "3",
          "status": "warning",
          "trend": "down",
          "trendValue": "-5"
        },
        {
          "id": "coverage",
          "type": "metric_card",
          "regionId": "kpi-row",
          "title": "Test Coverage",
          "value": "92%",
          "status": "good",
          "target": "90%"
        },
        {
          "id": "progress-chart",
          "type": "chart",
          "regionId": "main",
          "title": "Weekly Progress",
          "chartType": "line",
          "data": {
            "labels": ["W1", "W2", "W3", "W4", "W5", "W6"],
            "series": [
              {"name": "Actual", "values": [45, 55, 62, 70, 78, 85]},
              {"name": "Target", "values": [50, 60, 70, 75, 80, 85]}
            ]
          },
          "options": {
            "showLegend": true,
            "showGrid": true,
            "yAxisLabel": "Readiness %"
          }
        },
        {
          "id": "checklist",
          "type": "checklist",
          "regionId": "sidebar",
          "title": "Launch Checklist",
          "items": [
            {"text": "Feature freeze", "checked": true},
            {"text": "Security review", "checked": true},
            {"text": "Load testing", "checked": true},
            {"text": "Documentation", "checked": false},
            {"text": "Marketing ready", "checked": false}
          ]
        }
      ],
      "speakerNotes": "Key message: We're at 85% readiness, ahead of our 80% target."
    },
    {
      "id": "risks",
      "sectionId": "risks",
      "type": "content",
      "title": "Risks & Mitigations",
      "keyMessage": "3 risks identified, all with mitigation plans",
      "layout": {
        "template": "two_column"
      },
      "widgets": [
        {
          "id": "risk-list",
          "type": "risk_list",
          "regionId": "left",
          "title": "Open Risks",
          "items": [
            {
              "text": "API rate limits may impact launch",
              "severity": "high",
              "mitigation": "Implement caching layer",
              "owner": "Platform Team"
            },
            {
              "text": "Documentation incomplete",
              "severity": "medium",
              "mitigation": "Doc sprint this week",
              "owner": "Tech Writing"
            },
            {
              "text": "Partner integration pending",
              "severity": "low",
              "mitigation": "Fallback to manual process",
              "owner": "Partnerships"
            }
          ]
        },
        {
          "id": "decisions",
          "type": "decision_list",
          "regionId": "right",
          "title": "Pending Decisions",
          "items": [
            {
              "text": "Launch date confirmation",
              "description": "Confirm Feb 1 or delay to Feb 15",
              "owner": "VP Product",
              "deadline": "2025-01-20"
            },
            {
              "text": "Pricing tier structure",
              "description": "Finalize enterprise pricing",
              "owner": "Finance",
              "deadline": "2025-01-18"
            }
          ]
        }
      ],
      "speakerNotes": "Walk through each risk. Emphasize that all have mitigation plans."
    },
    {
      "id": "timeline",
      "sectionId": "next-steps",
      "type": "content",
      "title": "Launch Timeline",
      "widgets": [
        {
          "id": "timeline-table",
          "type": "table",
          "columns": [
            {"key": "date", "header": "Date", "width": "20%"},
            {"key": "milestone", "header": "Milestone", "width": "40%"},
            {"key": "owner", "header": "Owner", "width": "20%"},
            {"key": "status", "header": "Status", "width": "20%", "align": "center"}
          ],
          "rows": [
            {"date": "Jan 15", "milestone": "Feature freeze", "owner": "Engineering", "status": "Complete"},
            {"date": "Jan 20", "milestone": "Security audit", "owner": "Security", "status": "Complete"},
            {"date": "Jan 25", "milestone": "Load testing", "owner": "QA", "status": "In Progress"},
            {"date": "Jan 28", "milestone": "Documentation", "owner": "Tech Writing", "status": "Pending"},
            {"date": "Feb 1", "milestone": "Launch", "owner": "All Teams", "status": "Scheduled"}
          ]
        }
      ]
    },
    {
      "id": "summary",
      "type": "summary",
      "title": "Key Takeaways",
      "keyMessage": "On track for Feb 1 launch",
      "widgets": [
        {
          "id": "takeaways",
          "type": "content_block",
          "content": "## Summary\n\n1. **85% readiness** - Ahead of 80% target\n2. **3 risks identified** - All have mitigation plans\n3. **2 decisions needed** - By Jan 20\n\n**Next Review:** January 22, 2025",
          "format": "markdown"
        }
      ],
      "speakerNotes": "Emphasize the Feb 1 date and the two pending decisions."
    }
  ]
}
```

## Rendering

```bash
# HTML output
presspec render launch-readiness.json --output ./output

# With speaker notes
presspec render launch-readiness.json --output ./output --speaker-notes

# PDF output
presspec render launch-readiness.json --output ./output --format pdf

# Watch mode for development
presspec serve launch-readiness.json --port 8080
```

## Features Demonstrated

- **Metadata** - Title, description, author, tags
- **Theme** - Brand colors, logo, typography
- **Narrative** - Storyline and sections
- **Dashboard layout** - KPI row with content grid
- **Metric cards** - Values, trends, targets, sparklines
- **Charts** - Line chart with multiple series
- **Checklist** - Task completion status
- **Risk list** - Severity, mitigation, owner
- **Decision list** - Deadlines and owners
- **Table** - Structured data display
- **Speaker notes** - Presenter guidance
