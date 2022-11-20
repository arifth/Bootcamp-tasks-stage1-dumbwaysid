let data =[]
let datablog = []
let cardCtn = ""


function handleForm(e){
    
    // e.preventDefault

    // use obj desctructuring here 
    // https://www.freecodecamp.org/news/array-and-object-destructuring-in-javascript/
    
    let data = { name: "", start:"", end:"", desc:"", img:"" }

    let name = document.getElementById('name').value
    let start = document.getElementById('startdate').value
    let end = document.getElementById('enddate').value
    let desc = document.getElementById('desc').value
    let img = document.getElementById('input-img')

    let imgUrl = URL.createObjectURL(img.files[0])

    // console.log(typeof imgUrl)

    // console.log( name, start, end, desc, img)

   //assigning data to placeholder obj 
    let placeholder = {
        name: name,
        start: start,
        end: end,
        desc: desc,
        imgUrl: imgUrl,
    }

    console.log(placeholder)

    datablog.push(placeholder);
     
    console.log(datablog);
    renderCard()

}

function renderCard(){
    let card = ''
    
    let cardCtn = document.getElementById('section2')

    console.log(typeof datablog)
    
    datablog.forEach(handleRender);

    function handleRender(elem,index,array){
        console.log()
      card += `<div class="cardcont">
      <div class="img-cont">
        <img src="${elem.imgUrl}" alt="">
        <h2 class="judul">
          ${elem.name}
        </h2>
        <h3>
        ${durationHandle()}
        </h3>
        <p>
          ${elem.desc}
        </p>
        <div class="icon-cont">
          <i class="fa-brands fa-node-js fa-4x"></i>
          <i class="fa-brands fa-react fa-4x"></i>
          <i class="fa-brands fa-java fa-4x"></i>
          <i class="fa-solid fa-scroll fa-4x"></i>

        </div>
        <div class="button-cont">
          <button class="btn-scnd">delete </button>
          <button class="btn-scnd">edit</button>
        </div>
      </div>
    </div>`
    
    //render data to html , caution dangerous 
    console.log(cardCtn)
    cardCtn.innerHTML = card 

    }    

}

function durationHandle(){
    return "durasi 6 minggu"
}