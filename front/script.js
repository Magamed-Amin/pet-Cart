
const login = document.getElementById("login");
const register = document.getElementById("register");
const message = document.getElementById("message");

// Переключение на регистрацию
document.getElementById("toRegister").addEventListener("click", function() {
    login.classList.add("hidden");
    register.classList.remove("hidden");
    message.textContent = "";
});

// Переключение на авторизацию
document.getElementById("toLogin").addEventListener("click", function() {
    register.classList.add("hidden");
    login.classList.remove("hidden");
    message.textContent = "";
});

// Отправка формы авторизации
document.getElementById("loginForm").addEventListener("submit", function(e) {
    e.preventDefault();

    message.textContent = "Go";
});

// Отправка формы регистрации
document.getElementById("registerForm").addEventListener("submit", function(e) {
    e.preventDefault();

    message.textContent = "Go";
});
