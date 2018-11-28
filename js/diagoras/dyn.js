var constdata;
var state = 0;

$(document).ready(function() {
    $("#dynform").submit(function(e){
        e.preventDefault(e);
        $("#result").html("Παρακαλώ περιμένετε");
        var date = $("#datehidden").val();
        var label = $("#label").val();
        $.ajax({
            url: "/getdyn",
            data: {
                date:date,
                label:label,
            },
            accepts: {
                text: "application/json; charset=utf-8"
            },
            error: function() {
                $("#result").html("Σφάλμα κατά την ανάκτηση");
              },
            success: function(data){
                constdata = data
                parseDataMin(constdata)
            },
            timeout: 10000 // sets timeout to 10 seconds
        });
    });
});

function parseDataFull(data){
    $("#result").html("")
    $("#result").append("<h3>Δυναμολόγιο</h3>")
    $("#result").append("<button id=\"datatoggle\" class=\"btn btn-primary\">Συνοπτική</button>")
    $("#result").append("<h4>Παρόντες</h4>")
    if(constdata.proswpiko == null){
        $("#result").append("<h4>Κανείς παρών</h4>")    
    }else{
        unique_ranks = getUniqueValuesOfKey(constdata.proswpiko,"rank")
        for(var rank in unique_ranks){
            var temp = data.proswpiko.filter(obj => {
                return obj.rank === unique_ranks[rank]
            })
            $("#result").append("<h5>"+unique_ranks[rank]+"</h5>")
            ulelement = document.createElement('ul')
            for (var i in temp){
                lielement = document.createElement('li')
                lielement.style = "list-style:none"
                span1 = document.createElement('span')
                a = document.createElement('a')
                a.setAttribute('href',"retrieveproswpiko?id="+temp[i].id);
                a.innerHTML = temp[i].surname+" "+temp[i].name;
                span1.appendChild(a)
                lielement.appendChild(span1)
                ulelement.appendChild(lielement)
            }
            $("#result").append(ulelement)
        }
    }
    if (constdata.metaboles == null){
        $("#result").append("<h4>Καμία μεταβολή</h4>")
    }else{
        $("#result").append("<h4>Μεταβολές</h4>")
        unique_categories = getUniqueValuesOfKey(constdata.metaboles,"Category")
        for (var category in unique_categories){
            var temp = data.metaboles.filter(obj => {
                return obj.Category === unique_categories[category]
            })
            $("#result").append("<h5>"+unique_categories[category]+"</h5>")
            ulelement = document.createElement('ul')
            for (var i in temp){
                lielement = document.createElement('li')
                lielement.style = "list-style:none"
                span1 = document.createElement('span')
                a = document.createElement('a')
                a.setAttribute('href',"retrieveproswpiko?id="+temp[i].PersonID);
                a.innerHTML = temp[i].Surname+" "+temp[i].Name+" ";
                span1.appendChild(a)
                span2 = document.createElement('span')
                span2.innerHTML = temp[i].Repr
                lielement.appendChild(span1)
                lielement.appendChild(span2)
                ulelement.appendChild(lielement)
            }
            $("#result").append(ulelement)
        } 
      }
    if (constdata.ypiresies == null){
        $("#result").append("<h4>Καμία υπηρεσία</h4>")
    }else{
        $("#result").append("<h4>Υπηρεσίες</h4>")
        ulelement = document.createElement('ul')
        for(var ypiresia in constdata.ypiresies){
            lielement = document.createElement('li')
            lielement.style = "list-style:none"
            span1 = document.createElement('span')
            a = document.createElement('a')
            a.setAttribute('href',"retrieveproswpiko?id="+constdata.ypiresies[ypiresia].idperson);
            a.innerHTML = constdata.ypiresies[ypiresia].surname+" "+constdata.ypiresies[ypiresia].name;
            span1.appendChild(a)
            span2 = document.createElement('span')
            span2.innerHTML = " "+constdata.ypiresies[ypiresia].rank+" "+constdata.ypiresies[ypiresia].perigrafi
            lielement.appendChild(span1)
            lielement.appendChild(span2)
            ulelement.appendChild(lielement)
        }
        $("#result").append(ulelement)
    }
    if (constdata.aitiseis == null){
        $("#result").append("<h4>Καμία αίτηση</h4>")
    }else{
        $("#result").append("<h4>Αιτήσεις</h4>")
        ulelement = document.createElement('ul')
        for(var aitisi in constdata.aitiseis){
            lielement = document.createElement('li')
            lielement.style = "list-style:none"
            span1 = document.createElement('span')
            a = document.createElement('a')
            a.setAttribute('href',"retrieveproswpiko?id="+constdata.aitiseis[aitisi].idperson);
            a.innerHTML = constdata.aitiseis[aitisi].surname+" "+constdata.aitiseis[aitisi].name;
            span1.appendChild(a)
            span2 = document.createElement('span')
            span2.innerHTML = " "+constdata.aitiseis[aitisi].perigrafi
            lielement.appendChild(span1)
            lielement.appendChild(span2)
            ulelement.appendChild(lielement)
        }
        $("#result").append(ulelement)
    }
    if(constdata.anafores == null){
        $("#result").append("<h4>Καμία αναφορά</h4>")
    }else{
        $("#result").append("<h4>Αναφορές</h4>")
        ulelement = document.createElement('ul')
        for(var anafora in constdata.anafores){
            lielement = document.createElement('li')
            lielement.style = "list-style:none"
            span1 = document.createElement('span')
            a = document.createElement('a')
            a.setAttribute('href',"retrieveproswpiko?id="+constdata.anafores[anafora].idperson);
            a.innerHTML = constdata.anafores[anafora].surname+" "+constdata.anafores[anafora].name;
            span1.appendChild(a)
            span2 = document.createElement('span')
            span2.innerHTML = " "+constdata.anafora[anafores].perigrafi
            lielement.appendChild(span1)
            lielement.appendChild(span2)
            ulelement.appendChild(lielement)
        }
        $("#result").append(ulelement)
    }
    if(constdata.ergasies == null){
        $("#result").append("<h4>Καμία εργασία</h4>")
    }else{
        $("#result").append("<h4>Εργασίες</h4>")
        ulelement = document.createElement('ul')
        for(var ergasia in constdata.ergasies){
            lielement = document.createElement('li')
            lielement.style = "list-style:none"
            span1 = document.createElement('span')
            a = document.createElement('a')
            a.setAttribute('href',"retrieveproswpiko?id="+constdata.ergasies[ergasia].idperson);
            a.innerHTML = constdata.ergasies[ergasia].surname+" "+constdata.ergasies[ergasia].name;
            span1.appendChild(a)
            span2 = document.createElement('span')
            span2.innerHTML = " "+constdata.ergasies[ergasia].perigrafi
            lielement.appendChild(span1)
            lielement.appendChild(span2)
            ulelement.appendChild(lielement)
        }
        $("#result").append(ulelement)
    }
    $("#datatoggle").click(function(e){
        if (state == 0){
            state = 1
            parseDataFull(constdata)
        }else{
            state = 0
            parseDataMin(constdata)
        }
    });
};

function parseDataMin(data){
    $("#result").html("")
    $("#result").append("<h3>Δυναμολόγιο</h3>")
    $("#result").append("<button id=\"datatoggle\" class=\"btn btn-primary\">Αναλυτική</button>")
    if (constdata.proswpiko == null){
        $("#result").append("<h4>Κανείς παρών</h4>")
    }else{
        $("#result").append("<h4>Παρόντες</h4>")
        unique_ranks = getUniqueValuesOfKey(data.proswpiko,"rank")
        for (var index in unique_ranks){
            var temp = constdata.proswpiko.filter(obj => {
                return obj.rank === unique_ranks[index]
              })
            $("#result").append("<h5>"+unique_ranks[index]+" "+temp.length+"</h5>")
            }
    }
    if(data.metaboles == null){
        $("#result").append("<h4>Καμία μεταβολή</h4>")
    }else{
        $("#result").append("<h4>Μεταβολές</h4>")
        unique_categories = getUniqueValuesOfKey(data.metaboles,"Category")
        for (var index in unique_categories){
            var temp = data.metaboles.filter(obj => {
                return obj.Category === unique_categories[index]
              })
              $("#result").append("<h5>"+unique_categories[index]+" "+temp.length+"</h5>")
        }
    }
    if(constdata.ypiresies == null){
        $("#result").append("<h4>Καμία υπηρεσία</h4>")
    }else{
        $("#result").append("<h4>Υπηρεσίες "+constdata.ypiresies.length+"</h4>")
    }
    if(constdata.aitiseis == null){
        $("#result").append("<h4>Καμία αίτηση</h4>")
    }else{
        $("#result").append("<h4>Αιτήσεις "+constdata.aitiseis.length+"</h4>")
    }
    if(constdata.anafores == null){
        $("#result").append("<h4>Καμία αναφορά</h4>")
    }else{
        $("#result").append("<h4>Αναφορές"+constdata.anafores.length+"</h4>")
    }
    if(constdata.ergasies == null){
        $("#result").append("<h4>Καμία εργασία</h4>")
    }else{
        $("#result").append("<h4>Εργασίες "+constdata.ergasies.length+"</h4>")
    }
    $("#datatoggle").click(function(e){
        if (state == 0){
            state = 1
            parseDataFull(constdata)
        }else{
            state = 0
            parseDataMin(constdata)
        }
    });
}

function getUniqueValuesOfKey(array, key){
    return array.reduce(function(carry, item){
      if(item[key] && !~carry.indexOf(item[key])) carry.push(item[key]);
      return carry;
    }, []);
  }