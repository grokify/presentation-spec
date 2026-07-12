# Customizing Themes

This guide covers creating and customizing presentation themes.

## Overview

Themes control the visual appearance of presentations:

- Colors
- Typography
- Branding
- Layout density
- Aspect ratio

## Basic Theme

```json
{
  "theme": {
    "name": "my-theme",
    "aspectRatio": "16:9"
  }
}
```

## Colors

### Color Palette

```json
{
  "theme": {
    "colors": {
      "primary": "#1a73e8",
      "secondary": "#34a853",
      "accent": "#ea4335",
      "background": "#ffffff",
      "surface": "#f8f9fa",
      "text": "#1e293b",
      "textMuted": "#64748b"
    }
  }
}
```

| Color | Usage |
|-------|-------|
| `primary` | Main brand color, buttons, links |
| `secondary` | Secondary actions, accents |
| `accent` | Highlights, call-to-action |
| `background` | Page background |
| `surface` | Card and widget backgrounds |
| `text` | Primary text |
| `textMuted` | Secondary text, captions |

### Status Colors

```json
{
  "colors": {
    "success": "#22c55e",
    "warning": "#f59e0b",
    "error": "#ef4444"
  }
}
```

These are used by metric cards and status indicators.

## Typography

```json
{
  "theme": {
    "typography": {
      "fontFamily": "Inter, system-ui, sans-serif",
      "headingFontFamily": "Inter, system-ui, sans-serif",
      "codeFontFamily": "JetBrains Mono, monospace",
      "baseFontSize": "16px",
      "lineHeight": "1.6"
    }
  }
}
```

| Property | Description |
|----------|-------------|
| `fontFamily` | Body text font stack |
| `headingFontFamily` | Heading font stack |
| `codeFontFamily` | Code block font |
| `baseFontSize` | Base font size |
| `lineHeight` | Base line height |

## Branding

Add company branding:

```json
{
  "theme": {
    "brand": {
      "logo": "/assets/logo.png",
      "logoPosition": "top-right",
      "companyName": "Acme Corp",
      "primaryColor": "#1a73e8",
      "secondaryColor": "#34a853"
    }
  }
}
```

### Logo Positions

| Position | Description |
|----------|-------------|
| `top-left` | Upper left corner |
| `top-right` | Upper right corner |
| `bottom-left` | Lower left corner |
| `bottom-right` | Lower right corner |

## Density

Control spacing and content density:

```json
{
  "theme": {
    "density": "normal"
  }
}
```

| Density | Use Case |
|---------|----------|
| `compact` | More content, tighter spacing |
| `normal` | Balanced (default) |
| `spacious` | More whitespace, fewer elements |

## Aspect Ratio

```json
{
  "theme": {
    "aspectRatio": "16:9"
  }
}
```

| Ratio | Use Case |
|-------|----------|
| `16:9` | Widescreen displays (default) |
| `4:3` | Traditional screens |
| `1:1` | Square format |

## Dark Mode Theme

```json
{
  "theme": {
    "name": "dark-mode",
    "colors": {
      "primary": "#60a5fa",
      "secondary": "#4ade80",
      "accent": "#f472b6",
      "background": "#0f172a",
      "surface": "#1e293b",
      "text": "#f1f5f9",
      "textMuted": "#94a3b8",
      "success": "#22c55e",
      "warning": "#fbbf24",
      "error": "#f87171"
    },
    "typography": {
      "fontFamily": "Inter, sans-serif",
      "codeFontFamily": "Fira Code, monospace"
    }
  }
}
```

## Corporate Theme

```json
{
  "theme": {
    "name": "corporate",
    "density": "normal",
    "aspectRatio": "16:9",
    "brand": {
      "logo": "/assets/company-logo.png",
      "logoPosition": "top-right",
      "companyName": "Acme Corp"
    },
    "colors": {
      "primary": "#0f4c81",
      "secondary": "#2e7d32",
      "accent": "#ff6f00",
      "background": "#ffffff",
      "surface": "#f5f5f5",
      "text": "#212121",
      "textMuted": "#757575"
    },
    "typography": {
      "fontFamily": "Roboto, sans-serif",
      "headingFontFamily": "Roboto Slab, serif",
      "baseFontSize": "18px"
    }
  }
}
```

## Minimal Theme

```json
{
  "theme": {
    "name": "minimal",
    "density": "spacious",
    "colors": {
      "primary": "#000000",
      "secondary": "#666666",
      "background": "#ffffff",
      "surface": "#fafafa",
      "text": "#1a1a1a",
      "textMuted": "#888888"
    },
    "typography": {
      "fontFamily": "system-ui, sans-serif",
      "baseFontSize": "20px",
      "lineHeight": "1.8"
    }
  }
}
```

## Theme Inheritance

Override specific properties while keeping defaults:

```json
{
  "theme": {
    "name": "custom",
    "colors": {
      "primary": "#your-brand-color"
    }
  }
}
```

Only specified properties are overridden.

## Best Practices

1. **Use consistent colors** - Match your brand guidelines
2. **Test readability** - Ensure sufficient contrast
3. **Consider accessibility** - Use WCAG-compliant color combinations
4. **Use web-safe fonts** - Or include font files
5. **Test on target displays** - Check colors on projectors/screens
