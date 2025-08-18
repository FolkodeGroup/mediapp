"use client";
import Login from "../../components/auth/Login";
import MediappLogo from "../../assets/images/MediappLogo.png"
/* import FormBG from "../../assets/images/form-bg.png" */

const LoginView = () => {
  return (
    <section className="section-login">
        <img src={MediappLogo} alt="Mediapp Logo" className="image-login" />
        {/* <img src={FormBG} alt="Form BG" className=""/> */}
        <Login/>
    </section>
    );
};
export default LoginView;