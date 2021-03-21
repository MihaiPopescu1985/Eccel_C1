function showAddProjectModal() {
    let modal = document.getElementById("addProject")
    let view = window.getComputedStyle(modal).getPropertyValue('display')

    if (view == 'none') {
        modal.style.display = 'block'
    }
}

function closeAddProjectModal() {
    document.getElementById('new-project-form').reset()
    document.getElementById('addProject').style.display = 'none'
}