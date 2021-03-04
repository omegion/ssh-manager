package bw

import (
	"encoding/base64"
	"encoding/json"
)

// Folder is a folder in Bitwarden.
type Folder struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Field is custom fields under Item.
type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Item is a secure item of Bitwarden.
type Item struct {
	ID        string  `json:"id,omitempty"`
	Type      int     `json:"type"`
	Name      string  `json:"name"`
	FolderID  string  `json:"folderId"`
	Notes     []Field `json:"notes"`
	UpdatedAt string  `json:"revisionDate"`
}

// IsExists checks if Item is set.
func (i Item) IsExists() bool {
	return i.ID != ""
}

// Encode encodes Item for storing to Bitwarden.
func (i *Item) Encode() (string, error) {
	var p []byte

	p, err := json.Marshal(i)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(p), nil
}

// UnmarshalJSON is custom unmarshal for Item.
func (i *Item) UnmarshalJSON(data []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value["notes"] != nil {
		rawNotes := value["notes"].(string)

		decodedRawNotes, err := base64.StdEncoding.DecodeString(rawNotes)
		if err != nil {
			return err
		}

		var notes []Field

		err = json.Unmarshal(decodedRawNotes, &notes)
		if err != nil {
			return err
		}

		i.ID = value["id"].(string)
		i.Type = int(value["type"].(float64))
		i.Name = value["name"].(string)
		i.FolderID = value["folderId"].(string)
		i.Notes = notes
	}

	return nil
}

// MarshalJSON is custom marshal for Item.
func (i Item) MarshalJSON() ([]byte, error) {
	var notes []byte

	notes, err := json.Marshal(i.Notes)
	if err != nil {
		return nil, err
	}

	encodedNotes := base64.StdEncoding.EncodeToString(notes)

	return json.Marshal(map[string]interface{}{
		"name":     i.Name,
		"login":    i.Name,
		"type":     i.Type,
		"folderId": i.FolderID,
		"notes":    encodedNotes,
	})
}
