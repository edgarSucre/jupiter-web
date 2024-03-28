package components

// css classes for components
const (
	ButtonBase = "cursor-pointer rounded-lg border border-primary bg-primary font-semibold text-white transition hover:bg-opacity-90"
	LabelBase  = "mb-2.5 block font-medium text-black dark:text-white"
	//TextInputBase = "w-full rounded-lg border border-stroke bg-transparent font-medium py-4 pl-6 pr-10 outline-none focus:border-primary focus-visible:shadow-none dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
)

func TextInputBase(err string) string {
	base := "w-full rounded-lg border bg-transparent font-medium py-4 pl-6 pr-10 outline-none focus:border-primary focus-visible:shadow-none dark:bg-form-input dark:focus:border-primary"

	if err == "" {
		return base + " border-stroke dark:border-form-strokedark"
	}

	return base + " border-danger"
}
