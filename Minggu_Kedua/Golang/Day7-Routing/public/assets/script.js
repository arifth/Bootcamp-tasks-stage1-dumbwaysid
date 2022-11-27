// deklarasi variabel dari dom 
// Name, email, phone number, subject, your message ,
// masukkan dalam Object
// send email via mailto:




function handleSubmit(e){
   
    console.log('function tertekan')
    var nama = document.getElementById('jeneng').value 
    var email = document.getElementById('email').value 
    var phone= document.getElementById('phone').value 
    var subject = document.getElementById('subject').value
    var message = document.getElementById('message').value
    console.log('tertekan')
    console.log( nama, email, phone, subject, message );
    

    // declare new HTML node pointing to mail app
    let link = document.createElement('a')
    

    // inject it with element from form 
    link.href = `mailto:${email}?subject=${subject}&body=${message}`

    // simulate clicking href tag 
    link.click()
}