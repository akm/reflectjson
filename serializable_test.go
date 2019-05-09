package typedict

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/akm/typedict/typedict_test/foo"
)

func TestSerializableTypes(t *testing.T) {
	dict := New([]interface{}{
		(*foo.A)(nil),
	})

	dataTypes := SerializableTypes(dict.Types())
	if len(dataTypes) != 5 {
		t.Fatalf("length of dataTypes must be 5 but was %d", len(dataTypes))
	}

	{
		dt := dataTypes.DetectByName("A")
		assert.NotNil(t, dt)
		assert.Equal(t, "github.com/akm/typedict/typedict_test/foo", dt.PkgPath)
		assert.Equal(t, "struct", dt.Kind)
		assert.Nil(t, dt.Elem)
	}

	{
		dt := dataTypes.DetectByName("B")
		assert.NotNil(t, dt)
		assert.Equal(t, "github.com/akm/typedict/typedict_test/bar", dt.PkgPath)
		assert.Equal(t, "struct", dt.Kind)
		assert.Nil(t, dt.Elem)
	}

	{
		dt := dataTypes.DetectByName("C")
		assert.NotNil(t, dt)
		assert.Equal(t, "github.com/akm/typedict/typedict_test/baz", dt.PkgPath)
		assert.Equal(t, "struct", dt.Kind)
		assert.Nil(t, dt.Elem)
	}

	{
		d1 := dataTypes.DetectByName("D")
		{
			dt := d1
			assert.NotNil(t, dt)
			assert.Equal(t, "D", dt.Name)
			assert.Equal(t, "github.com/akm/typedict/typedict_test/baz", dt.PkgPath)
			assert.Equal(t, "slice", dt.Kind)
			assert.Equal(t, "baz.D", dt.Representation)
		}
		if assert.NotNil(t, d1.Elem) {
			d2 := d1.Elem
			{
				dt := d2
				assert.Equal(t, "", dt.Name)
				assert.Equal(t, "", dt.PkgPath)
				assert.Equal(t, "ptr", dt.Kind)
				assert.Equal(t, "*baz.C", dt.Representation)
			}
			if assert.NotNil(t, d2.Elem) {
				d3 := d2.Elem
				{
					dt := d3
					assert.Equal(t, "C", dt.Name)
					assert.Equal(t, "github.com/akm/typedict/typedict_test/baz", dt.PkgPath)
					assert.Equal(t, "struct", dt.Kind)
					assert.Equal(t, "baz.C", dt.Representation)
				}
			}
		}
	}
}
