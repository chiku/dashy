package app

import "strings"

type Interest struct {
	Name        string
	DisplayName string
}

func NewInterest(rawName string) Interest {
	parts := strings.Split(rawName, ":>")
	name, displayName := parts[0], parts[0]
	if len(parts) >= 2 && parts[1] != "" {
		displayName = parts[1]
	}

	return Interest{
		Name:        name,
		DisplayName: displayName,
	}
}

func (interest Interest) PipelineName(name string) (bool, string) {
	if !strings.EqualFold(name, interest.Name) || name == "" {
		return false, ""
	}
	return true, interest.DisplayName
}

type Interests struct {
	interests []Interest
}

func NewInterests() *Interests {
	return &Interests{}
}

func (interests *Interests) Add(name string) *Interests {
	interests.interests = append(interests.interests, NewInterest(name))
	return interests
}

func (interests Interests) PipelineName(name string) (int, string) {
	for i, interest := range interests.interests {
		match, displayName := interest.PipelineName(name)
		if match {
			return i, displayName
		}
	}

	return -1, ""
}
