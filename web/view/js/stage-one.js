function showMonthlyReportTable() {
    let table = document.getElementById("monthly-report-table")
    let tableStyles = window.getComputedStyle(table)
    let tableVisibility = tableStyles.getPropertyValue('display')

    tableVisibility == 'block' ? table.style.display = 'none' : table.style.display = 'block'
}