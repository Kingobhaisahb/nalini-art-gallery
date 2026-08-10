import { useEffect, useRef, useState } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);

  const googleButtonRef = useRef(null);
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();

    setLoading(true);
    setMessage("");

    try {
      const response = await axios.post(
        "http://localhost:8080/api/auth/login",
        {
          email: email,
          password: password,
        }
      );

      localStorage.setItem("token", response.data.token);

      setMessage("Login successful!");

      navigate("/profile");
    } catch (error) {
      if (error.response) {
        setMessage(error.response.data.message || "Login failed");
      } else {
        setMessage("Could not connect to server");
      }
    } finally {
      setLoading(false);
    }
  };

  const handleGoogleLogin = async (response) => {
  setMessage("Signing in with Google...");

  try {
    const result = await axios.post(
      "http://localhost:8080/api/auth/google",
      {
        id_token: response.credential,
      }
    );

    localStorage.setItem("token", result.data.token);

    setMessage("Google login successful!");

    navigate("/profile");
  } catch (error) {
    console.error(error);

    if (error.response) {
      setMessage(
        error.response.data.message || "Google login failed"
      );
    } else {
      setMessage("Could not connect to server");
    }
  }
};

useEffect(() => {
  const initializeGoogle = () => {
    if (!window.google || !googleButtonRef.current) {
      return;
    }

    window.google.accounts.id.initialize({
      client_id: import.meta.env.VITE_GOOGLE_CLIENT_ID,
      callback: handleGoogleLogin,
    });

    window.google.accounts.id.renderButton(
      googleButtonRef.current,
      {
        theme: "outline",
        size: "large",
        text: "continue_with",
        shape: "rectangular",
        width: 250,
      }
    );
  };

  const interval = setInterval(() => {
    if (window.google) {
      initializeGoogle();
      clearInterval(interval);
    }
  }, 100);

  return () => clearInterval(interval);
}, []);

  return (
    <div>
      <h1>Nalini Art Gallery</h1>

      <h2>Login</h2>

      <form onSubmit={handleLogin}>
        <div>
          <label>Email</label>

          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>

        <br />

        <div>
          <label>Password</label>

          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>

        <br />

        <button type="submit" disabled={loading}>
          {loading ? "Logging in..." : "Login"}
        </button>
      </form>

      <br />

      <div ref={googleButtonRef}></div>

      <p>{message}</p>

      <p>
        Don't have an account?{" "}
        <Link to="/signup">Create account</Link>
      </p>
    </div>
  );
}

export default Login;