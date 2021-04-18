package chiffrement

func Dechiffrer(motChiffre, codeSecret string) string {
	motClair := ""
	lgSecret := len(codeSecret)
	j := 0
	for i := 0; i < len(motChiffre); i++ {
		motClair += string((motChiffre[i] - codeSecret[j] + 26) %  26 + 'A')
		j = (j + 1) % lgSecret
	}
	return motClair
}