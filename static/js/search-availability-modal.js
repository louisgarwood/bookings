let availabilityBtn = document.getElementById("check-availability-button");
availabilityBtn.addEventListener("click", function(){
    let html = `
    <form id="check-availability-form" action="post" novalidate class="needs-validation">
        <div class="row" style="width: 99%" id="reservation-dates-modal">
            <div class="col">
                <input class="form-control" type="text" name="start" id="start" required placeholder="Arrival Date">
            </div>
            <div class="col">
                <input class="form-control" type="text" name="end" id="end" required placeholder="Departure">  
            </div>
        </div>
    </form>
    `

    attention.custom({
        title: "Check Availability", 
        msg: html, 
        callback: function(result){

            let form = document.getElementById("check-availability-form");
            let formData = new FormData(form);
            formData.append("csrf_token", "{{.CSRFToken}}");
            console.log(formData.get("csrf_token"));
            fetch("/search-availability-json", {
                method: "post",
                body: formData,
            })
                .then(response => response.json())
                .then(data => {
                    console.log(data);
                    console.log(data.ok);
                    console.log(data.message);
                });
            }});
});