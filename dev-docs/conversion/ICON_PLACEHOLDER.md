# Icon Placeholder

An **icon.png** file (256x256 pixels) is required for the Home Assistant add-on store.

## Requirements

- **Size**: 256x256 pixels
- **Format**: PNG with transparency support
- **Theme**: Weather-related (station, cloud, sun, etc.)
- **Style**: Simple, clear, recognizable at small sizes

## Suggested Icons

You can use a royalty-free icon from one of these sources:

### Option 1: Noun Project (Free with attribution)
- https://thenounproject.com/search/icons/?q=weather+station
- Download a weather station icon
- Resize to 256x256 pixels

### Option 2: Flaticon (Free with attribution)
- https://www.flaticon.com/search?word=weather%20station
- Download PNG at 256x256 size

### Option 3: Material Design Icons
- https://pictogrammers.com/library/mdi/
- Search for "weather" or "weather-partly-cloudy"
- Export as PNG at 256x256

### Option 4: Font Awesome (Free tier)
- https://fontawesome.com/search?q=weather&o=r
- Use icons like `cloud-sun`, `temperature-half`, or `wind`

### Option 5: Create Your Own
Use a tool like:
- Canva (https://www.canva.com/)
- GIMP (https://www.gimp.org/)
- Inkscape (https://inkscape.org/)

## Quick Command to Create Test Icon

If you have ImageMagick installed, you can create a simple placeholder:

```bash
convert -size 256x256 xc:skyblue \
  -font DejaVu-Sans-Bold -pointsize 40 -fill white \
  -gravity center -annotate +0+0 'VEVOR\nWeather' \
  icon.png
```

## Once You Have the Icon

1. Save it as `icon.png` in the project root directory
2. Ensure it's 256x256 pixels
3. Delete this `ICON_PLACEHOLDER.md` file
4. Commit the icon to your repository

## Optional: Add a Logo

A `logo.png` (also 256x256) is optional and will be displayed in the add-on documentation.
