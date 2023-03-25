package underscore

import "strings"


const A = 'A'
const Z = 'Z'

func ToCamelCase(str string) string {
    var stringBuilder strings.Builder

    for _, character := range str {
        if character >= A && character <= Z {
            stringBuilder.WriteRune('-')
        }
        stringBuilder.WriteRune(character)
    }

    return strings.ToLower(stringBuilder.String())
}
