let data =[]
let datablog = []
let cardCtn = ""
let check = []
let checkHolder = '' 


function handleForm(e){
    
    // e.preventDefault

    // use obj desctructuring here 
    // https://www.freecodecamp.org/news/array-and-object-destructuring-in-javascript/
    
   let data = { name: "", start:"", end:"", desc:"", img:"", checkHolder:"" }

    let name = document.getElementById('name').value
    let start = document.getElementById('startdate').value
    let end = document.getElementById('enddate').value
    let desc = document.getElementById('desc').value
    let img = document.getElementById('input-img')
    let check = document.querySelectorAll('input[type="checkbox"]')

    let imgUrl = URL.createObjectURL(img.files[0])
    
    check.forEach((elem) => {
    //check apakah cheklist di klik 
    if(elem.checked){
      checkHolder +=`<i class="fa-brands fa-${elem.value} fa-4x"></i>`
      
    }
    });
   
  console.log(checkHolder);
    // console.log( name, start, end, desc, img)

   //assigning data to placeholder obj 
    let placeholder = {
        name: name,
        start: start,
        end: end,
        desc: desc,
        imgUrl: imgUrl,
        checkHolder: checkHolder,
    }


    datablog.push(placeholder);
     
    renderCard()

    if(checkHolder!==""){
      checkHolder=""
    }

    console.log(placeholder)
    durationHandle(start, end)
   
}

function renderCard(){

    let card = ''
    
    let cardCtn = document.getElementById('section2')

    console.log(typeof datablog)
    
    datablog.forEach(handleRender);

    function handleRender(elem,index,array){
        console.log()
      card += `<div class="cardcont" >
      <a href="https://arifth.github.io/Bootcamp-tasks-stage1-dumbwaysid/Minggu_Pertama/Pengenalan_HTML/Day5DateFnStaticSite/card.html">
        <div class="img-cont">
          <img src="${elem.imgUrl}" alt="">
            <h2 class="judul">
              ${elem.name}
            </h2>
           <h3>
            Durasi Kursus ${durationHandle(elem.start, elem.end)} Hari
           </h3>
           <p>
            ${elem.desc}
           </p>
         <div class="icon-cont">
         ${checkHolder}
        </div>
        <div class="button-cont">
          <button class="btn-scnd">delete </button>
          <button class="btn-scnd">edit</button>
        </div>
      </div>
      </a>
    </div>`
    
    //render data to html , caution dangerous 
    //https://www.javascripttutorial.net/javascript-dom/javascript-innerhtml-vs-createelement/
    console.log(cardCtn)

    cardCtn.innerHTML = card
    
    }    

}

function durationHandle(start , end){
    let selisihMs= new Date(end) - new Date(start)
    console.log(selisihMs)
    // total selisih dalam milisecond di convert ke hari
    let selisihHari = selisihMs / (1000*60*60*24)
    console.log(selisihHari)
    return selisihHari
}

