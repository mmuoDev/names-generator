package names_test

import (
	"github.com/mmuoDev/names-generator/names"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamesAreRetrievedFromFiles(t *testing.T) {
	tribes := []string{"igbo", "yoruba", "hausa"}

	//male
	for _, tribe := range tribes {
		r, err := names.GetNames(tribe, "male")
		t.Log("err", err)
		assert.NotEmpty(t, r)

	}

	//female
	for _, tribe := range tribes {
		r, _ := names.GetNames(tribe, "female")
		assert.NotEmpty(t, r)

	}
}

func TestIgboNamesAreGenerated(t *testing.T) {
	res, _ := names.GenerateRandomNames("igbo", "male", 2)
	assert.NotEmpty(t, res)
	assert.Equal(t, len(res), 2)
}

func TestYorubaNamesAreGenerated(t *testing.T) {
	res, _ := names.GenerateRandomNames("yoruba", "female", 2)
	assert.NotEmpty(t, res)
	assert.Equal(t, len(res), 2)
}

func TestHausaNamesAreGenerated(t *testing.T) {
	res, _ := names.GenerateRandomNames("hausa", "female", 2)
	assert.NotEmpty(t, res)
	assert.Equal(t, len(res), 2)
}