import { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

function Profile() {
  const [user, setUser] = useState(null);
  const [message, setMessage] = useState("");

  const navigate = useNavigate();

  useEffect(() => {
    const getProfile = async () => {
      const token = localStorage.getItem("token");

      if (!token) {
        navigate("/login");
        return;
      }

      try {
        const response = await axios.get(
          "http://localhost:8080/api/profile",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );

        setUser(response.data.data);
      } catch (error) {
        localStorage.removeItem("token");
        setMessage("Session expired. Please login again.");
        navigate("/login");
      }
    };

    getProfile();
  }, [navigate]);

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  return (
    <div>
      <h1>Profile</h1>

      {user && (
        <div>
          <p>
            <strong>ID:</strong> {user.id}
          </p>

          <p>
            <strong>Name:</strong> {user.name}
          </p>

          <p>
            <strong>Email:</strong> {user.email}
          </p>

          <p>
            <strong>Phone:</strong> {user.phone}
          </p>

          <p>
            <strong>Role:</strong> {user.role}
          </p>

          <p>
            <strong>Auth Provider:</strong> {user.auth_provider}
          </p>
        </div>
      )}

      <p>{message}</p>

      <button onClick={handleLogout}>
        Logout
      </button>
    </div>
  );
}

export default Profile;