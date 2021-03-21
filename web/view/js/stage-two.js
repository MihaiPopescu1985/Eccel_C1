function closeModal(form, modal) {
    document.getElementById(form).reset()
    document.getElementById(modal).style.display = 'none'
}

function showModal(element) {
    let modal = document.getElementById(element)
    let view = window.getComputedStyle(modal).getPropertyValue('display')

    if (view == 'none') {
        modal.style.display = 'block'
    }
}