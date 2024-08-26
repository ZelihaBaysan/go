package lasagna

// PreparationTime calculates the preparation time based on layers and average preparation time per layer.
func PreparationTime(layers []string, averagePre int) int {
	if averagePre == 0 {
		averagePre = 2
	}
	return len(layers) * averagePre
}

// Quantities calculates the quantities of noodles and sauce needed based on the layers provided.
func Quantities(layers []string) (int, float64) {
	noodles, sauce := 0, float64(0)
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient adds the secret ingredient from the friends' list to the end of the user's list.
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales the ingredient quantities for the given number of portions.
func ScaleRecipe(list []float64, portions int) []float64 {
	sonuc := make([]float64, len(list))
	for i := 0; i < len(list); i++ {
		sonuc[i] = list[i] * float64(portions) / 2
	}
	return sonuc
}
