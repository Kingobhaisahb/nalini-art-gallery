import { useState } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

function Signup() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [password, setPassword] = useState("");

  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  const handleSignup = async (e) => {
    e.preventDefault();

    setLoading(true);
    setMessage("");

    try {
      const response = await axios.post(
        "http://localhost:8080/api/auth/signup",
        {
          name: name,
          email: email,
          phone: phone,
          password: password,
        }
      );

      setMessage(response.data.message || "Signup successful");

      setTimeout(() => {
        navigate("/login");
      }, 1000);
    } catch (error) {
      if (error.response) {
        setMessage(error.response.data.message || "Signup failed");
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

      <h2>Create Account</h2>

      <form onSubmit={handleSignup}>
        <div>
          <label>Name</label>

          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>

        <br />

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
          <label>Phone</label>

          <input
            type="text"
            value={phone}
            onChange={(e) => setPhone(e.target.value)}
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
          {loading ? "Creating account..." : "Signup"}
        </button>
      </form>

      <p>{message}</p>

      <p>
        Already have an account?{" "}
        <Link to="/login">Login</Link>
      </p>
    </div>
  );
}

export default Signup;