package main

import (
    gcolor "github.com/gookit/color"
    "image/color"
    "regexp"
    "strconv"
    "strings"
)

type ColoredTextPart struct {
    Text  string
    Color color.RGBA
}

func colPrintln(textWithColorCodes string) {
    lines := strings.Split(textWithColorCodes, "\n")
    for _, line := range lines {
        textParts := parseColorCodedText(line)
        for index, textPart := range textParts {
            textColor := textPart.Color
            r := textColor.R
            g := textColor.G
            b := textColor.B
            rgbColor := gcolor.Rgb(r, g, b, false)
            if index == len(textParts)-1 {
                rgbColor.Print(textPart.Text + "\n")
            } else {
                rgbColor.Print(textPart.Text)
            }
        }
    }
}
func parseColorCodedText(draw string) []ColoredTextPart {
    if draw == "" {
        return []ColoredTextPart{{
            Text:  "",
            Color: White,
        }}
    }
    var textParts []ColoredTextPart
    // color codes look like this [:red] [:blue] [:green] [:yellow] [:white] [:24,54,222]
    regexPattern := `\[:([\w,]+)\]`
    rp, _ := regexp.Compile(regexPattern)
    // split the string into parts
    stringParts := rp.Split(draw, -1)
    if len(stringParts) == 1 {
        return []ColoredTextPart{{
            Text:  draw,
            Color: White,
        }}
    }
    if stringParts[0] == "" {
        stringParts = stringParts[1:]
    }
    // find the color codes
    colorCodes := rp.FindAllStringSubmatch(draw, -1)
    if len(colorCodes) < len(stringParts) {
        // prepend [:white]
        colorCodes = append([][]string{{":white", "white"}}, colorCodes...)
    }
    for i, textPart := range stringParts {
        matchingColor := toColor(colorCodes, i)
        textParts = append(textParts, ColoredTextPart{
            Text:  textPart,
            Color: matchingColor,
        })
    }
    return textParts
}
func toColor(matches [][]string, index int) color.RGBA {
    if len(matches) == 0 {
        return White
    }
    colorName := matches[index][1]
    if strings.ContainsRune(colorName, ',') {
        colorValues := strings.Split(colorName, ",")
        rVal, _ := strconv.Atoi(colorValues[0])
        gVal, _ := strconv.Atoi(colorValues[1])
        bVal, _ := strconv.Atoi(colorValues[2])
        return color.RGBA{R: uint8(rVal), G: uint8(gVal), B: uint8(bVal), A: 255}
    }
    colorForText := colorFromName(colorName)
    return colorForText
}

func colorFromName(name string) color.RGBA {
    switch name {
    case "red":
        return Red
    case "blue":
        return Blue
    case "green":
        return Green
    case "yellow":
        return Yellow
    case "darkText":
        return DarkTextColor
    case "lightText":
        return LightTextColor
    case "magic":
        return MagicHighlightColor
    case "rare":
        return RareHighlightColor
    case "set":
        return SetHighlightColor
    case "unique":
        return UniqueHighlightColor
    default:
        return White
    }
}
