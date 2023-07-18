<script>
    import "../../app.css"
    import Container from "../../Container.svelte";

    import logo from "$lib/images/ease-logo.svg"
    import telegram from "$lib/images/telegram.svg"
    import divider from "$lib/images/line.svg"
</script>


<section class="h-screen bg-ease-grey">
    <Container classes="h-screen flex justify-center items-center">
        <div class="auth bg-aqua bg-white rounded-[20px] ">
            <div class="auth__content h-full py-[40px] px-[40px]">
                <div class="auth__content-header text-center">
                    <div class="auth__content-header-logo flex justify-center">
                        <img src={logo} alt="ease-logo">
                    </div>
                    <div class="auth__content-header-title mt-[20px] text-[20px] font-l">
                        Wallet Authentication
                    </div>
                    <div class="auth__content-header-info mt-[10px]">
                        Enter your account information for authentication.
                    </div>
                    <div class="auth__content-header-link-telegram mt-[20px]">
                        <a href="/" class="telegram-logo py-[12px] flex justify-center items-center bg-[#459DDA] text-white rounded-[9px]">
                            <span class="mr-[10px]"><img src={telegram} alt="telegram"></span>
                            Telegram
                        </a>
                    </div>
                    <div class="auth__content-header-divider flex items-center justify-between my-[10px]">
                        <img src={divider} alt="divider">
                        <span>or</span>
                        <img src={divider} alt="divider">
                    </div>
                </div>
                <div class="auth__content-form">
                    <form class="flex flex-col" action="POST">
                        <div class="auth__content-form-email">
                            <label for="email" class="block mb-[5px]">E-mail Address</label>
                            <input type="text" name="email" id="input-email" class="w-full border border-[#E2E3E6] py-[10px] px-4 rounded-[9px]">
                        </div>
                        <div class="auth__content-form-email mt-[10px]">
                            <label for="email" class="block mb-[5px]">Password</label>
                            <input type="text" name="email" id="input-email" class="w-full border border-[#E2E3E6] py-[10px] px-4 rounded-[9px]">
                        </div>
                        <a href="/" class="my-[15px] text-center">Forgot password?</a>
                        <button class="py-[12px] bg-[#0E1218] text-white rounded-[9px]">Authorize</button>
                    </form>
                    <div class="auth__content-form-register mt-[15px]">
                        New to Ease Wallet? <a href="/" class="font-xl">Create your account now</a>
                    </div>
                </div>
            </div>
        </div>
    </Container>
</section>