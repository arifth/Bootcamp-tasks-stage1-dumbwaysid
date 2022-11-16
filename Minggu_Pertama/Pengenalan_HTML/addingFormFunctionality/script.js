// deklarasi variabel dari dom 
// Name, email, phone number, subject, your message ,
// masukkan dalam Object
// send email via mailto:




function handleSubmit(e){

    var nama = document.getElementById('jeneng').value 
    var email = document.getElementById('email').value 
    var phone= document.getElementById('phone').value 
    var subject = document.getElementById('subject').value
    var message = document.getElementById('message').value
    console.log('tertekan')
    console.log( nama, email, phone, subject, message );
}