package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestFindAllRecipes(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		recipe := []string{"bread"}
		ingredients := [][]string{{"yeast", "flour"}}
		supplies := []string{"yeast", "flour", "corn"}
		assert.ElementsMatch(t, []string{"bread"}, leetcode.FindAllRecipes(recipe, ingredients, supplies))
	})

	t.Run("t2", func(t *testing.T) {
		recipe := []string{"bread", "sandwich"}
		ingredients := [][]string{{"yeast", "flour"}, {"bread", "meat"}}
		supplies := []string{"yeast", "flour", "meat"}
		assert.ElementsMatch(t, []string{"bread", "sandwich"}, leetcode.FindAllRecipes(recipe, ingredients, supplies))
	})

	t.Run("t3", func(t *testing.T) {
		recipe := []string{"bread", "burger", "sandwich"}
		ingredients := [][]string{{"yeast", "flour"}, {"sandwich", "meat", "bread"}, {"bread", "meat"}}
		supplies := []string{"yeast", "flour", "meat"}
		assert.ElementsMatch(t, []string{"bread", "sandwich", "burger"}, leetcode.FindAllRecipes(recipe, ingredients, supplies))
	})

	t.Run("t4", func(t *testing.T) {
		recipe := []string{"xevvq", "izcad", "p", "we", "bxgnm", "vpio", "i", "hjvu", "igi", "anp", "tokfq", "z", "kwdmb", "g", "qb", "q", "b", "hthy"}
		ingredients := [][]string{{"wbjr"}, {"otr", "fzr", "g"}, {"fzr", "wi", "otr", "xgp", "wbjr", "igi", "b"}, {"fzr", "xgp", "wi", "otr", "tokfq", "izcad", "igi", "xevvq", "i", "anp"}, {"wi", "xgp", "wbjr"}, {"wbjr", "bxgnm", "i", "b", "hjvu", "izcad", "igi", "z", "g"}, {"xgp", "otr", "wbjr"}, {"wbjr", "otr"}, {"wbjr", "otr", "fzr", "wi", "xgp", "hjvu", "tokfq", "z", "kwdmb"}, {"xgp", "wi", "wbjr", "bxgnm", "izcad", "p", "xevvq"}, {"bxgnm"}, {"wi", "fzr", "otr", "wbjr"}, {"wbjr", "wi", "fzr", "xgp", "otr", "g", "b", "p"}, {"otr", "fzr", "xgp", "wbjr"}, {"xgp", "wbjr", "q", "vpio", "tokfq", "we"}, {"wbjr", "wi", "xgp", "we"}, {"wbjr"}, {"wi"}}
		supplies := []string{"wi", "otr", "wbjr", "fzr", "xgp"}
		assert.ElementsMatch(t, []string{"xevvq", "izcad", "bxgnm", "i", "hjvu", "tokfq", "z", "g", "b", "hthy"}, leetcode.FindAllRecipes(recipe, ingredients, supplies))
	})

	t.Run("t5", func(t *testing.T) {
		recipe := []string{"bread"}
		ingredients := [][]string{{"yeast", "flour"}}
		supplies := []string{"yeast", "corn"}
		assert.ElementsMatch(t, []string(nil), leetcode.FindAllRecipes(recipe, ingredients, supplies))
	})
}

func TestFindAllRecipesV2(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		recipe := []string{"bread"}
		ingredients := [][]string{{"yeast", "flour"}}
		supplies := []string{"yeast", "flour", "corn"}
		assert.ElementsMatch(t, []string{"bread"}, leetcode.FindAllRecipesV2(recipe, ingredients, supplies))
	})

	t.Run("t2", func(t *testing.T) {
		recipe := []string{"bread", "sandwich"}
		ingredients := [][]string{{"yeast", "flour"}, {"bread", "meat"}}
		supplies := []string{"yeast", "flour", "meat"}
		assert.ElementsMatch(t, []string{"bread", "sandwich"}, leetcode.FindAllRecipesV2(recipe, ingredients, supplies))
	})

	t.Run("t3", func(t *testing.T) {
		recipe := []string{"bread", "burger", "sandwich"}
		ingredients := [][]string{{"yeast", "flour"}, {"sandwich", "meat", "bread"}, {"bread", "meat"}}
		supplies := []string{"yeast", "flour", "meat"}
		assert.ElementsMatch(t, []string{"bread", "sandwich", "burger"}, leetcode.FindAllRecipesV2(recipe, ingredients, supplies))
	})

	t.Run("t4", func(t *testing.T) {
		recipe := []string{"xevvq", "izcad", "p", "we", "bxgnm", "vpio", "i", "hjvu", "igi", "anp", "tokfq", "z", "kwdmb", "g", "qb", "q", "b", "hthy"}
		ingredients := [][]string{{"wbjr"}, {"otr", "fzr", "g"}, {"fzr", "wi", "otr", "xgp", "wbjr", "igi", "b"}, {"fzr", "xgp", "wi", "otr", "tokfq", "izcad", "igi", "xevvq", "i", "anp"}, {"wi", "xgp", "wbjr"}, {"wbjr", "bxgnm", "i", "b", "hjvu", "izcad", "igi", "z", "g"}, {"xgp", "otr", "wbjr"}, {"wbjr", "otr"}, {"wbjr", "otr", "fzr", "wi", "xgp", "hjvu", "tokfq", "z", "kwdmb"}, {"xgp", "wi", "wbjr", "bxgnm", "izcad", "p", "xevvq"}, {"bxgnm"}, {"wi", "fzr", "otr", "wbjr"}, {"wbjr", "wi", "fzr", "xgp", "otr", "g", "b", "p"}, {"otr", "fzr", "xgp", "wbjr"}, {"xgp", "wbjr", "q", "vpio", "tokfq", "we"}, {"wbjr", "wi", "xgp", "we"}, {"wbjr"}, {"wi"}}
		supplies := []string{"wi", "otr", "wbjr", "fzr", "xgp"}
		assert.ElementsMatch(t, []string{"xevvq", "izcad", "bxgnm", "i", "hjvu", "tokfq", "z", "g", "b", "hthy"}, leetcode.FindAllRecipesV2(recipe, ingredients, supplies))
	})

	t.Run("t5", func(t *testing.T) {
		recipe := []string{"bread"}
		ingredients := [][]string{{"yeast", "flour"}}
		supplies := []string{"yeast", "corn"}
		assert.ElementsMatch(t, []string(nil), leetcode.FindAllRecipesV2(recipe, ingredients, supplies))
	})
}
