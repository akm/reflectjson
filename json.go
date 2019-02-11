package reflectjson

import (
	"encoding/json"
	"io"
)

func Output(dest io.Writer, obj interface{}) error {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}
	{
		_, err := dest.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
