package util

func ValidateMember(name string) (bool, string) {

    if len(name) > 50 {
        return false, "your input name is too long"
    }

    return true, ""
}
