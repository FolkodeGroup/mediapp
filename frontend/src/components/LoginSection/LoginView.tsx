"use client";
import Login from "../../components/auth/Login";
import MediappLogo from "../../assets/images/MediappLogo.png"

const LoginView = () => {
  return (
    <section>
  <img src={MediappLogo} alt="Mediapp Logo" style={{ width: 120, margin: "0 auto" }} />
        <Login/>
    </section>
    );
};
export default LoginView;