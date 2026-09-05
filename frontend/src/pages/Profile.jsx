import { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

function Profile() {
  const [user, setUser] = useState(null);
  const [message, setMessage] = useState("");

  // Payment states
  const [addressId, setAddressId] = useState("");
  const [paymentLoading, setPaymentLoading] = useState(false);

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

  const handlePayment = async () => {
    if (!addressId) {
      setMessage("Please enter your address ID.");
      return;
    }

    const token = localStorage.getItem("token");

    if (!token) {
      setMessage("Please login again.");
      navigate("/login");
      return;
    }

    setPaymentLoading(true);
    setMessage("");

    try {
      // Step 1:
      // Create our order + Razorpay order
      const response = await axios.post(
        "http://localhost:8080/api/orders/checkout",
        {
          address_id: Number(addressId),
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      const order = response.data.order;
      const payment = response.data.payment;

      // Step 2:
      // Configure Razorpay Checkout
      const options = {
        key: import.meta.env.VITE_RAZORPAY_KEY_ID,

        amount: Math.round(payment.amount * 100),

        currency: payment.currency,

        name: "Nalini Art Gallery",

        description: `Payment for Order #${order.id}`,

        order_id: payment.razorpay_order_id,

        prefill: {
          name: user?.name || "",
          email: user?.email || "",
          contact: user?.phone || "",
        },

        theme: {
          color: "#3399cc",
        },

        // Step 3:
        // Razorpay calls this after successful payment
        handler: async function (razorpayResponse) {
          try {
            setMessage("Payment successful. Verifying payment...");

            // Step 4:
            // Send Razorpay payment details to our backend
            const verifyResponse = await axios.post(
              "http://localhost:8080/api/payments/verify",
              {
                order_id: order.id,
                razorpay_order_id:
                  razorpayResponse.razorpay_order_id,
                razorpay_payment_id:
                  razorpayResponse.razorpay_payment_id,
                razorpay_signature:
                  razorpayResponse.razorpay_signature,
              },
              {
                headers: {
                  Authorization: `Bearer ${token}`,
                },
              }
            );

            setMessage(
              verifyResponse.data.message ||
                "Payment verified successfully!"
            );
          } catch (error) {
            console.error("Payment verification error:", error);

            if (error.response) {
              setMessage(
                error.response.data.message ||
                  "Payment verification failed."
              );
            } else {
              setMessage(
                "Payment was completed, but verification failed."
              );
            }
          }
        },

        // Called if user closes Razorpay Checkout
        modal: {
          ondismiss: function () {
            setMessage("Payment window closed.");
          },
        },
      };

      // Step 5:
      // Open Razorpay Checkout
      if (!window.Razorpay) {
        setMessage(
          "Razorpay Checkout failed to load. Please refresh the page."
        );
        return;
      }

      const razorpay = new window.Razorpay(options);

      razorpay.on("payment.failed", function (response) {
        console.error("Razorpay payment failed:", response);

        setMessage(
          response.error?.description ||
            "Payment failed. Please try again."
        );
      });

      razorpay.open();
    } catch (error) {
      console.error("Checkout error:", error);

      if (error.response) {
        setMessage(
          error.response.data.message ||
            "Unable to start checkout."
        );
      } else {
        setMessage("Could not connect to server.");
      }
    } finally {
      setPaymentLoading(false);
    }
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
            <strong>Auth Provider:</strong>{" "}
            {user.auth_provider}
          </p>
        </div>
      )}

      <hr />

      <h2>Test Payment</h2>

      <div>
        <label>
          Address ID:
        </label>

        <br />

        <input
          type="number"
          value={addressId}
          onChange={(e) => setAddressId(e.target.value)}
          placeholder="Enter address ID"
        />
      </div>

      <br />

      <button
        onClick={handlePayment}
        disabled={paymentLoading}
      >
        {paymentLoading
          ? "Starting Payment..."
          : "Pay with Razorpay"}
      </button>

      <p>{message}</p>

      <hr />

      <button onClick={handleLogout}>
        Logout
      </button>
    </div>
  );
}

export default Profile;