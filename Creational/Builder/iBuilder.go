package builder

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func GetBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}

	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
