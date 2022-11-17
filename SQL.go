// Biron bir tablitsa uchun yangi column qo'shilganda, avvalgi ma'lumotlarni yangilash uchun :
// Bizda firstname va lastname columnla boridi. biza fullname column qo'shdik. oldingi memborlani fullnameni qilish uchun mashi kod yozildi: (FARKHOD)
UPDATE author SET fullname=firstname || ' ' || lastname || ' ' || middlename;

