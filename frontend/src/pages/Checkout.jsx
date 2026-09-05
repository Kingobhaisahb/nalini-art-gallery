import { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

function Checkout() {
  const [addressId, setAddressId] = useState("");
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const navigate = useNavigate();

  const handleCheckout = async () => {
    const token = localStorage.getItem("token");

    if (!token) {
      navigate("/login");
      return;
    }

    if (!addressId) {
      setMessage("Please enter your address ID");
      return;
    }

    setLoading(true);
    setMessage("");

    try {
      // Step 1:
      // Ask our backend to create the order
      // and Razorpay order.
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
      // Open Razorpay Checkout.
      const options = {
        key: import.meta.env.VITE_RAZORPAY_KEY_ID,

        amount: Math.round(payment.amount * 100),

        currency: payment.currency,

        name: "Nalini Art Gallery",

        description: `Payment for Order #${order.id}`,

        order_id: payment.razorpay_order_id,

        prefill: {
          name: "",
          email: "",
          contact: "",
        },

        handler: async function (razorpayResponse) {
          try {
            // Step 3:
            // Send Razorpay's payment details
            // to our backend for verification.
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
                "Payment successful!"
            );

          } catch (error) {
            console.error(error);

            if (error.response) {
              setMessage(
                error.response.data.message ||
                  "Payment verification failed"
              );
            } else {
              setMessage(
                "Could not verify payment with server"
              );
            }
          }
        },

        modal: {
          ondismiss: function () {
            setMessage("Payment cancelled");
          },
        },

        theme: {
          color: "#000000",
        },
      };

      const razorpay = new window.Razorpay(options);

      razorpay.on("payment.failed", function (response) {
        console.error("Payment failed:", response);

        setMessage(
          response.error?.description ||
            "Payment failed"
        );
      });

      razorpay.open();

    } catch (error) {
      console.error(error);

      if (error.response) {
        setMessage(
          error.response.data.message ||
            "Checkout failed"
        );
      } else {
        setMessage(
          "Could not connect to server"
        );
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <h1>Nalini Art Gallery</h1>

      <h2>Checkout</h2>

      <div>
        <label>Address ID</label>

        <input
          type="number"
          value={addressId}
          onChange={(e) => setAddressId(e.target.value)}
          placeholder="Enter address ID"
        />
      </div>

      <br />

      <button
        onClick={handleCheckout}
        disabled={loading}
      >
        {loading ? "Processing..." : "Proceed to Payment"}
      </button>

      <p>{message}</p>
    </div>
  );
}

export default Checkout;