package filepaths


// :Path at last represents the complete path to the file from the root directory
// :Name at last represents only the file name with its extension
type CompanyPaths struct {

	CompanyDashboardPath string

	NewJobFormTemplatePath string 
	NewTestFormTemplateName string 

	CompanyMyProfileTemplatePath string 
	CompanyMyApplicantsTemplatePath string
	CompanyMyJobListingsTemplatePath string
	CompanyScheduledTemplatePath string
	CompanyCompletedTemplatePath string
	FeedbacksTemplatePath string
}


// Direct (the file, not basepaths) paths to all Dashboards
func LoadCompanyPaths() CompanyPaths {
	return CompanyPaths{
		CompanyDashboardPath: "./template/dashboard/companydashboard.html",

		NewJobFormTemplatePath: "./template/company/newjobform.html",
		NewTestFormTemplateName: "newtest.html",

		CompanyMyProfileTemplatePath: "./template/company/myProfile.html",
		CompanyMyApplicantsTemplatePath: "./template/company/myapplicants.html",
		CompanyMyJobListingsTemplatePath: "./template/company/myjoblistings.html",
		CompanyScheduledTemplatePath: "./template/company/scheduled.html",
		CompanyCompletedTemplatePath: "./template/company/completed.html",
		FeedbacksTemplatePath: "./template/company/feedbacks.html",
	}
}
