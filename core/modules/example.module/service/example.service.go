package example_service

import "licor_model/core/server/shared"

func GetService() {
	shared.GetDB().QueryRow("query")
}
