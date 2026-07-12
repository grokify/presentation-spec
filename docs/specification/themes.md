# Themes

Themes control the visual styling of presentations.

## Structure

```json
{
  "theme": {
    "name": "corporate",
    "density": "normal",
    "aspectRatio": "16:9",
    "brand": { ... },
    "colors": { ... },
    "typography": { ... }
  }
}
```

## Fields

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Theme name |
| `density` | string | `compact`, `normal`, `spacious` |
| `aspectRatio` | string | `16:9`, `4:3`, `1:1` |
| `brand` | Brand | Branding options |
| `colors` | ColorPalette | Color scheme |
| `typography` | Typography | Font settings |

## Density

| Value | Description |
|-------|-------------|
| `compact` | Tighter spacing, more content |
| `normal` | Balanced spacing (default) |
| `spacious` | More whitespace |

## Aspect Ratio

| Value | Use Case |
|-------|----------|
| `16:9` | Widescreen (default) |
| `4:3` | Traditional screens |
| `1:1` | Square format |

## Brand

```json
{
  "brand": {
    "logo": "/assets/logo.png",
    "logoPosition": "top-right",
    "companyName": "Acme Corp",
    "primaryColor": "#1a73e8",
    "secondaryColor": "#34a853"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `logo` | string | Logo image path |
| `logoPosition` | string | `top-left`, `top-right`, `bottom-left`, `bottom-right` |
| `companyName` | string | Company name |
| `primaryColor` | string | Primary brand color (hex) |
| `secondaryColor` | string | Secondary brand color (hex) |

## Colors

```json
{
  "colors": {
    "primary": "#1a73e8",
    "secondary": "#34a853",
    "accent": "#ea4335",
    "background": "#ffffff",
    "surface": "#f8f9fa",
    "text": "#1e293b",
    "textMuted": "#64748b",
    "success": "#22c55e",
    "warning": "#f59e0b",
    "error": "#ef4444"
  }
}
```

| Field | Description |
|-------|-------------|
| `primary` | Primary UI color |
| `secondary` | Secondary UI color |
| `accent` | Accent color |
| `background` | Page background |
| `surface` | Card/widget background |
| `text` | Primary text color |
| `textMuted` | Secondary text color |
| `success` | Success indicator |
| `warning` | Warning indicator |
| `error` | Error indicator |

## Typography

```json
{
  "typography": {
    "fontFamily": "Inter, system-ui, sans-serif",
    "headingFontFamily": "Inter, system-ui, sans-serif",
    "codeFontFamily": "JetBrains Mono, monospace",
    "baseFontSize": "16px",
    "lineHeight": "1.6"
  }
}
```

| Field | Description |
|-------|-------------|
| `fontFamily` | Body text font |
| `headingFontFamily` | Heading font |
| `codeFontFamily` | Code block font |
| `baseFontSize` | Base font size |
| `lineHeight` | Base line height |

## Example

```json
{
  "theme": {
    "name": "dark-mode",
    "density": "normal",
    "aspectRatio": "16:9",
    "colors": {
      "primary": "#60a5fa",
      "background": "#0f172a",
      "surface": "#1e293b",
      "text": "#f1f5f9",
      "textMuted": "#94a3b8"
    },
    "typography": {
      "fontFamily": "Inter, sans-serif",
      "codeFontFamily": "Fira Code, monospace"
    }
  }
}
```
