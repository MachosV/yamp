$(document).ready(aitiseiscounter());
$(document).ready(getmaxaitisi());
$(document).ready(addlistenerloadmore());
$(document).ready(signaitisi());
$(document).ready(getcsrftoken());
var max;
var end = false;
var count
var csrftoken
table = document.getElementById("aitiseisTable")

function getcsrftoken(){
    csrftoken = $("#csrftoken").val();
}

function signaitisi(){
    $(".sign").off()
    $(".sign").click(function(e){
        this.blur()
        e.preventDefault()
        btn = this
        $.ajax({
            type: "POST",
            url: "/signaitisi",
            data: {
                csrftoken: csrftoken,
                sign: btn.classList[3],
                id: btn.parentElement.parentElement.dataset.idaitisi
            },
            error: function(){
                console.log("error signing")
            },
            success: function(data){
                $(btn.parentElement.parentElement).fadeOut(300, function() {})
            }
        })
    })
}

function addlistenerloadmore(){
    $("#loadmore").click(function(e){
        document.getElementById("loadmore").blur()
        e.preventDefault()
        if(end){return}
        document.getElementById("loadmorecontent").innerHTML = "Παρακαλώ περιμένετε"
        $.ajax({
            url: "/getaitiseis",
            data: {
                maxid: max
            },
            accepts: {
                text: "application/json; charset=utf-8"
            },
            error: function() {
                document.getElementById("loadmorecontent").innerHTML = "Σφάλμα κατά την ανάκτηση"
              },
            success: function(data){
                for(row in data){
                    rowtable = table.insertRow(-1)
                    rowtable.classList.add("aitisi")
                    rowtable.setAttribute('data-idaitisi',data[row].id)
                    perigrafi = rowtable.insertCell(-1);
                    bathmos = rowtable.insertCell(-1);
                    epitheto = rowtable.insertCell(-1);
                    onoma = rowtable.insertCell(-1)
                    date = rowtable.insertCell(-1)
                    monada = rowtable.insertCell(-1)
                    egkrisi = rowtable.insertCell(-1)
                    aporripsi = rowtable.insertCell(-1)
                    perigrafi.innerHTML = "<a href=/raitisi?id="+data[row].id+">"+data[row].perigrafi+"</a>"
                    bathmos.innerHTML = data[row].rank
                    epitheto.innerHTML = data[row].surname
                    onoma.innerHTML = data[row].name
                    date.innerHTML = data[row].date
                    monada.innerHTML = data[row].monada
                    egkrisi.innerHTML = "<a href=\"#\" role=\"button\" class=\"btn btn-success sign 1\">Έγκριση <span class=\"glyphicon glyphicon-ok\"></span></a>"
                    aporripsi.innerHTML = "<a href=\"#\" role=\"button\" class=\"btn btn-danger sign 0\">Απόρριψη <span class=\"glyphicon glyphicon-remove\"></span></a>"
                }
                if(data.length > 3){
                    document.getElementById("loadmorecontent").innerHTML = "Περισσότερες (<span id=\"counter\">"+count+"</span>)"
                }else{
                    document.getElementById("loadmorecontent").innerHTML = "Τέλος αιτήσεων"
                    end = true;
                }
                getmaxaitisi();
                aitiseiscounter();
                signaitisi();
            },
            timeout: 10000 // sets timeout to 10 seconds
        });
    })
}

function getmaxaitisi(){
    aitiseis = document.getElementsByClassName("aitisi")
    nums = [].slice.call(aitiseis).map(function(e){return e.dataset.idaitisi})    
    max = Math.max.apply(null,nums);
}

function aitiseiscounter(){
    try{
        spel = document.getElementById("counter")
        spel.innerHTML = spel.innerHTML - 4
        count = spel.innerHTML
        if (count == '0'){
            console.log("True")
            document.getElementById("loadmorecontent").innerHTML = "Τέλος αιτήσεων"
            end = true;
        }
    }catch (err){
    }
}

