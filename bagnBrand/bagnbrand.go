package main

import "fmt"

func main() {
	data := []map[string]interface{}{
		{"Account_type": "Asset", "aid": "achead::7", "name": "Non Current Asset", "parent_id": "achead::2"},
		{"Account_type": "Asset", "aid": "achead::8", "name": "Machinery", "parent_id": "achead::7"},
		{"Account_type": "Asset", "aid": "achead::7", "name": "Non Current Asset", "parent_id": "achead::2"},
		{"Account_type": "Asset", "aid": "achead::8", "name": "Machinery", "parent_id": "achead::7"},
		{"Account_type": "Asset", "aid": "achead::9", "name": "Current Asset", "parent_id": "achead::2"},
		{"Account_type": "Asset", "aid": "achead::10", "name": "Account Receivable", "parent_id": "achead::9"},
		{"Account_type": "Asset", "aid": "achead::11", "name": "Advance Deposit & Prepayment", "parent_id": "achead::9"},
		{"Account_type": "Asset", "aid": "achead::12", "name": "Cash at Bank", "parent_id": "achead::9"},
		{"Account_type": "Asset", "aid": "achead::13", "name": "Cash in Hand", "parent_id": "achead::9"},
		{"Account_type": "Asset", "aid": "achead::14", "name": "Inventory", "parent_id": "achead::9"},
		{"Account_type": "Liability", "aid": "achead::15", "name": "Longterm Liability", "parent_id": "achead::3"},
		{"Account_type": "Liability", "aid": "achead::16", "name": "Short-term Liability", "parent_id": "achead::3"},
		{"Account_type": "Liability", "aid": "achead::17", "name": "Bank Loan", "parent_id": "achead::15"},
		{"Account_type": "Liability", "aid": "achead::18", "name": "Current Liability", "parent_id": "achead::3"},
		{"Account_type": "Liability", "aid": "achead::19", "name": "Account Payable", "parent_id": "achead::18"},
		{"Account_type": "Liability", "aid": "achead::20", "name": "Duties & Taxes", "parent_id": "achead::5"},
		{"Account_type": "Revenue", "aid": "achead::21", "name": "Sales Accout", "parent_id": "achead::5"},
		{"Account_type": "Revenue", "aid": "achead::22", "name": "Indirect Income", "parent_id": "achead::4"},
		{"Account_type": "Equity", "aid": "achead::23", "name": "Capital Account", "parent_id": "achead::23"},
		{"Account_type": "Equity", "aid": "achead::24", "name": "Drawing Account", "parent_id": "achead::4"},
		{"Account_type": "Equity", "aid": "achead::25", "name": "Retained Earning", "parent_id": "achead::4"},
		{"Account_type": "Equity", "aid": "achead::26", "name": "Opening Balance Equity", "parent_id": "achead::4"},
		{"Account_type": "Expense", "aid": "achead::27", "name": "Direct Expense parent", "parent_id": "achead::6"},
		{"Account_type": "Expense", "aid": "achead::28", "name": "Cost of Revenue", "parent_id": "achead::27"},
		{"Account_type": "Expense", "aid": "achead::29", "name": "COGS", "parent_id": "achead::27"},
		{"Account_type": "Expense", "aid": "achead::30", "name": "Indirect Expense", "parent_id": "achead::6"},
		{"Account_type": "Expense", "aid": "achead::31", "name": "Administrative Expense", "parent_id": "achead::30"},
		{"Account_type": "Expense", "aid": "achead::32", "name": "Salary & Allowance", "parent_id": "achead::31"},
		{"Account_type": "Expense", "aid": "achead::34", "name": "Cash Discount", "parent_id": "achead::30"},
	}

	output := listAccount(data)

	fmt.Println(output)
}

func listAccount(data []map[string]interface{}) (result []string) {

	var acType string
	fmt.Print("Enter Your Account Type: ")
	fmt.Scan(&acType)
	for _, item := range data {
		//fmt.Println(i, item["Account_type"])

		if item["Account_type"] == acType {
			val := fmt.Sprintf("%#v", item["name"])
			result = append(result, val)
		}
	}
	return
}

// 2d map
// package main

// import "fmt"

// func main() {
// 	data := map[int]map[string]string{
// 		0:  {"Account_type": "Asset", "aid": "achead::7", "name": "Non Current Asset", "parent_id": "achead::2"},
// 		1:  {"Account_type": "Asset", "aid": "achead::8", "name": "Machinery", "parent_id": "achead::7"},
// 		2:  {"Account_type": "Asset", "aid": "achead::9", "name": "Current Asset", "parent_id": "achead::2"},
// 		3:  {"Account_type": "Asset", "aid": "achead::10", "name": "Account Receivable", "parent_id": "achead::9"},
// 		4:  {"Account_type": "Asset", "aid": "achead::11", "name": "Advance Deposit & Prepayment", "parent_id": "achead::9"},
// 		5:  {"Account_type": "Asset", "aid": "achead::12", "name": "Cash at Bank", "parent_id": "achead::9"},
// 		6:  {"Account_type": "Asset", "aid": "achead::13", "name": "Cash in Hand", "parent_id": "achead::9"},
// 		7:  {"Account_type": "Asset", "aid": "achead::14", "name": "Inventory", "parent_id": "achead::9"},
// 		8:  {"Account_type": "Liability", "aid": "achead::15", "name": "Longterm Liability", "parent_id": "achead::3"},
// 		9:  {"Account_type": "Liability", "aid": "achead::16", "name": "Short-term Liability", "parent_id": "achead::3"},
// 		10: {"Account_type": "Liability", "aid": "achead::17", "name": "Bank Loan", "parent_id": "achead::15"},
// 		11: {"Account_type": "Liability", "aid": "achead::18", "name": "Current Liability", "parent_id": "achead::3"},
// 		12: {"Account_type": "Liability", "aid": "achead::19", "name": "Account Payable", "parent_id": "achead::18"},
// 		13: {"Account_type": "Liability", "aid": "achead::20", "name": "Duties & Taxes", "parent_id": "achead::5"},
// 		14: {"Account_type": "Revenue", "aid": "achead::21", "name": "Sales Accout", "parent_id": "achead::5"},
// 		15: {"Account_type": "Revenue", "aid": "achead::22", "name": "Indirect Income", "parent_id": "achead::4"},
// 		16: {"Account_type": "Equity", "aid": "achead::23", "name": "Capital Account", "parent_id": "achead::23"},
// 		17: {"Account_type": "Equity", "aid": "achead::24", "name": "Drawing Account", "parent_id": "achead::4"},
// 		18: {"Account_type": "Equity", "aid": "achead::25", "name": "Retained Earning", "parent_id": "achead::4"},
// 		19: {"Account_type": "Equity", "aid": "achead::26", "name": "Opening Balance Equity", "parent_id": "achead::4"},
// 		20: {"Account_type": "Expense", "aid": "achead::27", "name": "Direct Expense parent", "parent_id": "achead::6"},
// 		21: {"Account_type": "Expense", "aid": "achead::28", "name": "Cost of Revenue", "parent_id": "achead::27"},
// 		22: {"Account_type": "Expense", "aid": "achead::29", "name": "COGS", "parent_id": "achead::27"},
// 		23: {"Account_type": "Expense", "aid": "achead::30", "name": "Indirect Expense", "parent_id": "achead::6"},
// 		24: {"Account_type": "Expense", "aid": "achead::31", "name": "Administrative Expense", "parent_id": "achead::30"},
// 		25: {"Account_type": "Expense", "aid": "achead::32", "name": "Salary & Allowance", "parent_id": "achead::31"},
// 		26: {"Account_type": "Expense", "aid": "achead::34", "name": "Cash Discount", "parent_id": "achead::30"},
// 	}
// 	// data2 := make([]string, map)
// 	res := listAccount(data)

// 	fmt.Println("res:", res)

// }

// func listAccount(data map[int]map[string]string) (result []string) {
// 	sliceIndex := []int{}
// 	sliceStr := []string{}
// 	var str string
// 	fmt.Print("Enter your account type: ")
// 	fmt.Scanln(&str)

// 	for key, element := range data {
// 		for _, value := range element {
// 			if value == str {
// 				sliceIndex = append(sliceIndex, key)
// 				sliceStr = append(sliceStr, data[key]["name"])
// 			}
// 		}
// 	}
// 	smallest := sliceIndex[0]
// 	for _, num := range sliceIndex[1:] {
// 		if num < smallest {
// 			smallest = num
// 		}
// 	}
// 	length := len(sliceIndex)
// 	result = make([]string, length)

// 	for i := 0; i < length; i++ {
// 		result[sliceIndex[i]-smallest] = sliceStr[i]
// 	}

// 	return

// }
