import { useState } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);

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

      <button>
        Continue with Google
      </button>

      <p>{message}</p>

      <p>
        Don't have an account?{" "}
        <Link to="/signup">Create account</Link>
      </p>
    </div>
  );
}

export default Login;