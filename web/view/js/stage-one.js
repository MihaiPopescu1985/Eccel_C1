function changeTableView() {
    let detailedTable = document.getElementById("detailed-table");
    let standardTable = document.getElementById("standard-table");
    let viewType = document.getElementById("table-view-type");

    let tableStyles = window.getComputedStyle(standardTable);
    let tableVisibility = tableStyles.getPropertyValue('display');

    if (tableVisibility == 'table')  {
        standardTable.style.display = 'none';
        detailedTable.style.display = 'table';
        viewType.innerHTML = "Vizualizare detaliata";
    } else {
        standardTable.style.display = 'table';
        detailedTable.style.display = 'none';
        viewType.innerHTML = "Vizualizare standard";
    }
}