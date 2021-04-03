/*
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

function showModal() {
    let modal = document.getElementById("addWorkday")
    let view = window.getComputedStyle(modal).getPropertyValue('display')

    if (view == 'none') {
        modal.style.display = 'block'
    }
}

function closeModal() {
    document.getElementById('new-workday-form').reset()
    document.getElementById('addWorkday').style.display = 'none'
}
*/