

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
    
    // let inputCechk = check.forEach((elem) => {
    //check apakah cheklist di klik 
    // if(elem.checked){
    //   checkHolder +=`<i class="fa-brands fa-${elem.value} fa-4x"></i>`
      
    // }
    // });
   
  // console.log(inputCechk);
    // console.log( name, start, end, desc, img)

   //assigning data to placeholder obj 
    let placeholder = {
        name: name,
        start: start,
        end: end,
        desc: desc,
        imgUrl: imgUrl,
        // checkHolder: inputCechk,
    }

    console.log(placeholder)
   
  }
