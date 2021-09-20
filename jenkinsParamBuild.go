package main

import (
	"fmt"
	"html/template"
	"log"
	"net/url"
	"os"
	"strconv"

	"golang.org/x/net/html"
)

type SimpleMap map[string]string

func HumanReadableJenkinsParamBuildUrl(htmlBuildUrl string) error {
	buildUrl := html.UnescapeString(htmlBuildUrl)
	fmt.Println(buildUrl)
	parsedUrl, err := url.Parse(buildUrl)
	if err != nil {
		log.Fatal("unable to parse the url")
		return err
	}
	for k, v := range parsedUrl.Query() {
		//fmt.Printf("%s : %s,\n", strconv.Quote(k), strconv.Quote(v[0]))
		fmt.Println(k, v)
	}
	return nil
}

func GenerateJenkinsParamBuildUrl(path string, m *SimpleMap) string {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "core-ci.oxid-esales.com"
	if path == "" {
		u.Path = "job/compilation-trigger-manual/parambuild"
	} else {
		u.Path = path
	}
	q := u.Query()
	for k, v := range *m {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func createHtmlContent(url string, value string) {
	const tpl = `
<li>
    <a href='{{.Url}}'>{{.Value}}</a>
</li>`
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("jenkins").Parse(tpl)
	check(err)
	data := struct {
		Url   string
		Value string
	}{
		Url:   url,
		Value: value,
	}
	err = t.Execute(os.Stdout, data)
	check(err)
}

func main() {
	buildUrl := "https://core-ci.oxid-esales.com/job/compilation-trigger-manual/parambuild/?DISPLAYNAME=6.3.x&amp;CE_BRANCH_OR_TAG=b-6.3.x&amp;PE_COMPOSER_VERSION=dev-b-6.3.x&amp;EE_COMPOSER_VERSION=dev-b-6.3.x&amp;PHP_VERSION=7.4&amp;MYSQL_VERSION=mysql%3A5.7&amp;TESTING_LIBRARY_VERSION=dev-b-6.3.x&amp;AZURE_THEME_VERSION=v1.4.2&amp;OXIDESHOP_DEMODATA_VERSION=dev-b-6.0&amp;FLOW_THEME_VERSION=%22dev-b-3.x%20as%20v3.x.x-dev%22&amp;WAVE_THEME_VERSION=dev-b-1.x&amp;OXIDESHOP_COMPOSER_PLUGIN_VERSION=%22dev-b-6.x%20as%20v5.x.x-dev%22&amp;OXIDESHOP_UNIFIED_NAMESPACE_GENERATOR_VERSION=%22dev-b-6.3.x%20as%20v2.x.x-dev%22&amp;OXIDESHOP_DB_VIEWS_GENERATOR_VERSION=%5Ev1.1.1&amp;OXIDESHOP_DEMODATA_INSTALLER_VERSION=%22dev-b-6.3.x%20as%20v1.x.x-dev%22&amp;OXIDESHOP_DOCTRINE_MIGRATION_WRAPPER_VERSION=%22dev-b-6.3.x%20as%20v3.x.x-dev%22&amp;OXIDESHOP_FACTS_VERSION=%22dev-b-6.3.x%20as%20v2.x.x-dev%22&amp;AMAZONPAY_MODULE_VERSION=%22dev-OXDEV-4477_changes_for_php8%20as%203.6.8%22&amp;GDPR_OPTIN_MODULE_VERSION=%5Ev2.3.3&amp;KLARNA_MODULE_VERSION=%5Ev5.5.1&amp;PAYMORROW_MODULE_VERSION=%5Ev2.0.4&amp;PAYONE_MODULE_VERSION=%22dev-OXDEV-4477_changes_for_php8%20as%20v1.5.0%22&amp;PAYPAL_MODULE_VERSION=%5Ev6.3.0&amp;USERCENTRICS_MODULE_VERSION=%5Ev1.1.3&amp;WYSIWYG_EDITOR_MODULE_VERSION=%5Ev2.4.0&amp;VISUALCMS_MODULE_VERSION=%5Ev3.4.0&amp;TEST_ENVIRONMENT_CE=export%20ACTIVATE_ALL_MODULES%3D1%20RUN_TESTS_FOR_MODULES%3D0%20RUN_TESTS_FOR_SHOP%3D1%20TEST_DIRS%3D%22%22%20PARTIAL_MODULE_PATHS%3Dbestit%2Famazonpay4oxid%2Coe%2Foepaypal%2Cddoe%2Fwysiwyg%2Coxps%2Fpaymorrow%2Coe%2Fgdproptin%2Cfc%2Ffcpayone%2Ctc%2Ftcklarna%2Coxps%2Fusercentrics%20REMOVE_FROM_PHPUNIT%3D%22libpng.warning%22&amp;TEST_ENVIRONMENT_PE=export%20ACTIVATE_ALL_MODULES%3D1%20RUN_TESTS_FOR_MODULES%3D0%20RUN_TESTS_FOR_SHOP%3D1%20TEST_DIRS%3D%22%22%20PARTIAL_MODULE_PATHS%3Dbestit%2Famazonpay4oxid%2Coe%2Foepaypal%2Cddoe%2Fvisualcms%2Cddoe%2Fwysiwyg%2Coxps%2Fpaymorrow%2Coe%2Fgdproptin%2Cfc%2Ffcpayone%2Ctc%2Ftcklarna%2Coxps%2Fusercentrics%20ADDITIONAL_TEST_PATHS%3Dvendor%2Foxid-esales%2Foxideshop-pe%2FTests%20REMOVE_FROM_PHPUNIT%3D%22libpng.warning%22&amp;TEST_ENVIRONMENT_EE=export%20ACTIVATE_ALL_MODULES%3D1%20RUN_TESTS_FOR_MODULES%3D0%20RUN_TESTS_FOR_SHOP%3D1%20TEST_DIRS%3D%22%22%20PARTIAL_MODULE_PATHS%3Dbestit%2Famazonpay4oxid%2Coe%2Foepaypal%2Cddoe%2Fvisualcms%2Cddoe%2Fwysiwyg%2Coxps%2Fpaymorrow%2Coe%2Fgdproptin%2Cfc%2Ffcpayone%2Ctc%2Ftcklarna%2Coxps%2Fusercentrics%20ADDITIONAL_TEST_PATHS%3Dvendor%2Foxid-esales%2Foxideshop-ee%2FTests%2Cvendor%2Foxid-esales%2Foxideshop-pe%2FTests%20REMOVE_FROM_PHPUNIT%3D%22libpng.warning%22&amp;COMPOSER_VCS_REPOSITORIES=oxid-esales%2Foxideshop-pe%2Chttps%3A%2F%2Fgithub.com%2FOXID-eSales%2Foxideshop_pe%3Boxid-esales%2Foxideshop-demodata-pe%2Chttps%3A%2F%2Fgithub.com%2FOXID-eSales%2Foxideshop_demodata_pe.git%3Boxid-esales%2Foxideshop-ee%2Chttps%3A%2F%2Fgithub.com%2FOXID-eSales%2Foxideshop_ee%3Boxid-esales%2Foxideshop-demodata-ee%2Chttps%3A%2F%2Fgithub.com%2FOXID-eSales%2Foxideshop_demodata_ee.git%3Bddoe%2Fvisualcms-module%2Chttps%3A%2F%2Fgithub.com%2FOXID-eSales%2Fvisual_cms_module.git%3Bpayone-gmbh%2Foxid-6%2Chttps%3A%2F%2Fgithub.com%2FOXID-eSales%2Foxid-6.git%3Bbestit%2Famazonpay4oxid%2Chttps%3A%2F%2Fgithub.com%2FOXID-eSales%2Famazon-pay-oxid.git"
	if err := HumanReadableJenkinsParamBuildUrl(buildUrl); err != nil {
		return
	}
	values := SimpleMap{
		"AMAZONPAY_MODULE_VERSION":                      strconv.Quote("dev-OXDEV-4477_changes_for_php8 as 3.6.8"),
		"AZURE_THEME_VERSION":                           "v1.4.2",
		"CE_BRANCH_OR_TAG":                              "b-6.4.x",
		"COMPOSER_ADDITIONAL_PACKAGES":                  "symfony/expression-language:^4.4",
		"COMPOSER_VCS_REPOSITORIES":                     "oxid-esales/oxideshop-pe,https://github.com/OXID-eSales/oxideshop_pe;oxid-esales/oxideshop-demodata-pe,https://github.com/OXID-eSales/oxideshop_demodata_pe.git;oxid-esales/oxideshop-ee,https://github.com/OXID-eSales/oxideshop_ee;oxid-esales/oxideshop-demodata-ee,https://github.com/OXID-eSales/oxideshop_demodata_ee.git;ddoe/visualcms-module,https://github.com/OXID-eSales/visual_cms_module.git;payone-gmbh/oxid-6,https://github.com/OXID-eSales/oxid-6.git;bestit/amazonpay4oxid,https://github.com/OXID-eSales/amazon-pay-oxid.git",
		"DISPLAYNAME":                                   "6.4.x",
		"EE_COMPOSER_VERSION":                           "dev-b-6.4.x",
		"FLOW_THEME_VERSION":                            strconv.Quote("dev-b-3.x as v3.x.x-dev"),
		"GDPR_OPTIN_MODULE_VERSION":                     "^v2.3.3",
		"KLARNA_MODULE_VERSION":                         "^v5.5.1",
		"MYSQL_VERSION":                                 "mysql:5.7",
		"OXIDESHOP_COMPOSER_PLUGIN_VERSION":             "^v5.2.0",
		"OXIDESHOP_DB_VIEWS_GENERATOR_VERSION":          "^v1.1.1",
		"OXIDESHOP_DEMODATA_INSTALLER_VERSION":          "^v1.2.0",
		"OXIDESHOP_DEMODATA_VERSION":                    "v6.0.4",
		"OXIDESHOP_DOCTRINE_MIGRATION_WRAPPER_VERSION":  strconv.Quote("dev-b-6.4.x as v3.x.x-dev"),
		"OXIDESHOP_FACTS_VERSION":                       "^v2.4.0",
		"OXIDESHOP_UNIFIED_NAMESPACE_GENERATOR_VERSION": "^v2.2.0",
		"PAYMORROW_MODULE_VERSION":                      "^v2.0.4",
		"PAYONE_MODULE_VERSION":                         strconv.Quote("dev-OXDEV-4477_changes_for_php8 as v1.5.0"),
		"PAYPAL_MODULE_VERSION":                         "^v6.3.0",
		"PE_COMPOSER_VERSION":                           "dev-b-6.4.x",
		"PHP_VERSION":                                   "8.0",
		"TEST_ENVIRONMENT_CE":                           "export ACTIVATE_ALL_MODULES=1 RUN_TESTS_FOR_MODULES=0 RUN_TESTS_FOR_SHOP=1 TEST_DIRS=\"\" PARTIAL_MODULE_PATHS=bestit/amazonpay4oxid,oe/oepaypal,ddoe/wysiwyg,oxps/paymorrow,oe/gdproptin,fc/fcpayone,tc/tcklarna,oxps/usercentrics REMOVE_FROM_PHPUNIT=\"libpng.warning\"",
		"TEST_ENVIRONMENT_EE":                           "export ACTIVATE_ALL_MODULES=1 RUN_TESTS_FOR_MODULES=0 RUN_TESTS_FOR_SHOP=1 TEST_DIRS=\"\" PARTIAL_MODULE_PATHS=bestit/amazonpay4oxid,oe/oepaypal,ddoe/visualcms,ddoe/wysiwyg,oxps/paymorrow,oe/gdproptin,fc/fcpayone,tc/tcklarna,oxps/usercentrics ADDITIONAL_TEST_PATHS=vendor/oxid-esales/oxideshop-ee/Tests,vendor/oxid-esales/oxideshop-pe/Tests REMOVE_FROM_PHPUNIT=\"libpng.warning\"",
		"TEST_ENVIRONMENT_PE":                           "export ACTIVATE_ALL_MODULES=1 RUN_TESTS_FOR_MODULES=0 RUN_TESTS_FOR_SHOP=1 TEST_DIRS=\"\" PARTIAL_MODULE_PATHS=bestit/amazonpay4oxid,oe/oepaypal,ddoe/visualcms,ddoe/wysiwyg,oxps/paymorrow,oe/gdproptin,fc/fcpayone,tc/tcklarna,oxps/usercentrics ADDITIONAL_TEST_PATHS=vendor/oxid-esales/oxideshop-pe/Tests REMOVE_FROM_PHPUNIT=\"libpng.warning\"",
		"TESTING_LIBRARY_VERSION":                       "dev-b-6.4.x",
		"USERCENTRICS_MODULE_VERSION":                   "^v1.1.3",
		"VISUALCMS_MODULE_VERSION":                      "^v3.4.0",
		"WAVE_THEME_VERSION":                            "v1.6.1",
		"WYSIWYG_EDITOR_MODULE_VERSION":                 "^v2.4.0",
	}
	createHtmlContent(GenerateJenkinsParamBuildUrl("", &values), "6.4.x")
}
