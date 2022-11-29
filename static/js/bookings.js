let attention = Prompt();


function Prompt() {
    let toast = function (c) {
    const {
        msg = "",
        icon = "success",
        position = "top-end",
    } = c;

    const Toast = Swal.mixin({
        toast: true,
        title: msg,
        position: position,
        icon: icon,
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: (toast) => {
        toast.addEventListener('mouseenter', Swal.stopTimer)
        toast.addEventListener('mouseleave', Swal.resumeTimer)
        }
    })

    Toast.fire({})
    }

    let success = function (c) {
    const {
        msg = "",
        icon = "success",
        text = "Something went right!",
    } = c;
    Swal.fire({
        icon: icon,
        title: msg,
        text: text,
    })
    }
    
    async function custom(c){
        const {
            msg = "",
            title ="",

        } = c;

        const { value: result } = await Swal.fire({
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: true,
            showCancelButton: true,
            willOpen: () => {
                if(c.willOpen !== undefined){
                    c.willOpen();
                }
            },
            preConfirm: () => {
                if(c.preConfirm !== undefined){
                    c.willOpen();
                }
            }
        })

        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
                if (result.value !== "") {
                    if (c.callback !== undefined) {
                        c.callback(result);
                    }
                }else{
                    c.callback(false);
                }
            }else{
                c.callback(false);
            }
        }
    }

    return {
    toast: toast,
    success: success,
    custom: custom,
    }
}

function alert(alertType, message){
    notie.alert({
    type: alertType,
    text: message,
    position: "top",
    })
}

function sweetAlert(alertType, message){
    Swal.fire(
    'Good job!',
    message,
    alertType
    )
}

(() => {
    'use strict'
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')

    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
    form.addEventListener('submit', event => {
        if (!form.checkValidity()) {
        event.preventDefault()
        event.stopPropagation()
        }

        form.classList.add('was-validated')
    }, false)
    })
})()