import { Link } from "react-router-dom";
import "./Home.css";

const categories = [
  {
    title: "Landscapes",
    image:
      "https://images.unsplash.com/photo-1500530855697-b586d89ba3ee?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "Florals",
    image:
      "https://images.unsplash.com/photo-1490750967868-88aa4486c946?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "Abstract",
    image:
      "https://images.unsplash.com/photo-1549490349-8643362247b5?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "Figurative",
    image:
      "https://images.unsplash.com/photo-1579783902614-a3fb3927b6a5?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "Still Life",
    image:
      "https://images.unsplash.com/photo-1577083552431-6e5fd01988a5?auto=format&fit=crop&w=900&q=85",
  },
];

const featuredPaintings = [
  {
    title: "Monsoon Village",
    price: "₹32,000",
    image:
      "https://images.unsplash.com/photo-1500534623283-312aade485b7?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "Evening Serenity",
    price: "₹27,500",
    image:
      "https://images.unsplash.com/photo-1470770841072-f978cf4d019e?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "White Blossoms",
    price: "₹24,000",
    image:
      "https://images.unsplash.com/photo-1497250681960-ef046c08a56e?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "Himalayan Trails",
    price: "₹28,000",
    image:
      "https://images.unsplash.com/photo-1464822759023-fed622ff2c3b?auto=format&fit=crop&w=900&q=85",
  },
];

const latestPaintings = [
  {
    title: "Golden Reflections",
    price: "₹30,000",
    image:
      "https://images.unsplash.com/photo-1500534314209-a25ddb2bd429?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "Serene Mountains",
    price: "₹26,500",
    image:
      "https://images.unsplash.com/photo-1464278533981-50106e6176b1?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "Blooming Grace",
    price: "₹29,000",
    image:
      "https://images.unsplash.com/photo-1495231916356-a86217efff12?auto=format&fit=crop&w=900&q=85",
  },
  {
    title: "Urban Dreams",
    price: "₹25,000",
    image:
      "https://images.unsplash.com/photo-1518005020951-eccb494ad742?auto=format&fit=crop&w=900&q=85",
  },
];

function PaintingCard({ painting }) {
  return (
    <article className="painting-card">
      <div className="painting-image-wrapper">
        <img src={painting.image} alt={painting.title} />
        <button className="cart-button" aria-label={`Add ${painting.title} to cart`}>
          🛒
        </button>
      </div>

      <div className="painting-card-details">
        <h3>{painting.title}</h3>
        <p>{painting.price}</p>
      </div>
    </article>
  );
}

function Home() {
  return (
    <div className="home-page">
      <header className="site-header">
        <Link to="/" className="brand">
          <div className="brand-symbol">♧</div>

          <div>
            <h1>NALINI ART GALLERY</h1>
            <span>ART LIVES FOREVER</span>
          </div>
        </Link>

        <nav className="desktop-nav">
          <Link to="/" className="active">
            Home
          </Link>
          <Link to="/gallery">Gallery</Link>
          <Link to="/about">About</Link>
          <Link to="/contact">Contact</Link>
        </nav>

        <div className="header-actions">
          <div className="search-box">
            <span>⌕</span>
            <input type="text" placeholder="Search artworks..." />
          </div>

          <button className="header-icon" aria-label="Shopping cart">
            🛒
          </button>

          <Link to="/profile" className="header-icon" aria-label="Profile">
            ♙
          </Link>
        </div>
      </header>

      <main>
        <section className="hero-section">
          <img
            className="hero-background"
            src="https://images.unsplash.com/photo-1500534623283-312aade485b7?auto=format&fit=crop&w=2200&q=90"
            alt="Landscape painting"
          />

          <div className="hero-overlay"></div>

          <div className="hero-content">
            <p className="eyebrow">ORIGINAL ARTWORKS</p>

            <h2>
              Art That
              <br />
              Brings Life Home
            </h2>

            <p className="hero-description">
              Discover a collection of original paintings by Nalini, where
              every brushstroke tells a story and every canvas holds a piece
              of the artist's soul.
            </p>

            <Link to="/gallery" className="primary-button">
              Explore the Collection <span>→</span>
            </Link>
          </div>

          <p className="hero-quote">
            “In every colour, a feeling.
            <br />
            In every painting, a story.”
            <span>— NALINI</span>
          </p>
        </section>

        <section className="category-section section-container">
          <div className="category-grid">
            {categories.map((category) => (
              <Link to="/gallery" className="category-card" key={category.title}>
                <img src={category.image} alt={category.title} />
                <div className="category-overlay"></div>

                <div className="category-content">
                  <h3>{category.title}</h3>
                  <span>EXPLORE →</span>
                </div>
              </Link>
            ))}
          </div>
        </section>

        <section className="collection-section section-container">
          <div className="section-heading">
            <div>
              <h2>Featured Paintings</h2>
              <p>HANDPICKED ARTWORKS FROM THE COLLECTION</p>
            </div>

            <Link to="/gallery" className="text-link">
              View All Paintings →
            </Link>
          </div>

          <div className="painting-grid">
            {featuredPaintings.map((painting) => (
              <PaintingCard painting={painting} key={painting.title} />
            ))}
          </div>
        </section>

        <section className="artist-process-section">
          <div className="artist-image">
            <img
              src="https://images.unsplash.com/photo-1513364776144-60967b0f800f?auto=format&fit=crop&w=1200&q=85"
              alt="Artist painting on a canvas"
            />
          </div>

          <div className="artist-story">
            <p className="eyebrow dark-eyebrow">THE ARTIST</p>
            <h2>Nalini</h2>

            <p>
              For me, painting is more than art — it is a way of expressing
              emotions, capturing moments, and finding beauty in everyday
              life. My work is inspired by nature, people, and the world
              around me.
            </p>

            <Link to="/about" className="secondary-button">
              Read My Story →
            </Link>
          </div>

          <div className="process-story">
            <p className="eyebrow dark-eyebrow">BEHIND THE ART</p>
            <h2>The Process</h2>

            <p>
              From the first brushstroke to the final masterpiece, watch how
              the artwork comes alive.
            </p>

            <Link to="/process" className="text-link">
              Watch the Process →
            </Link>
          </div>

          <div className="process-image">
            <img
              src="https://images.unsplash.com/photo-1579783902614-a3fb3927b6a5?auto=format&fit=crop&w=1000&q=85"
              alt="Painting process with brushes and colours"
            />

            <button className="play-button" aria-label="Watch painting process">
              ▶
            </button>
          </div>
        </section>

        <section className="collection-section section-container">
          <div className="section-heading">
            <div>
              <h2>Latest Arrivals</h2>
              <p>FRESHLY ADDED TO THE GALLERY</p>
            </div>

            <Link to="/gallery" className="text-link">
              View All →
            </Link>
          </div>

          <div className="painting-grid">
            {latestPaintings.map((painting) => (
              <PaintingCard painting={painting} key={painting.title} />
            ))}
          </div>
        </section>

        <section className="contact-banner">
          <div>
            <h2>Looking for Something Special?</h2>
            <p>Get in touch for commissions or any inquiries.</p>
          </div>

          <Link to="/contact" className="primary-button">
            Contact Us →
          </Link>
        </section>
      </main>

      <footer className="site-footer">
        <div className="footer-brand">
          <h2>NALINI ART GALLERY</h2>
          <p>ART LIVES FOREVER</p>
          <span>Original paintings. Real stories. A more beautiful world.</span>

          <div className="social-links">
            <span>◎</span>
            <span>●</span>
            <span>▶</span>
          </div>
        </div>

        <div className="footer-column">
          <h3>Quick Links</h3>
          <Link to="/">Home</Link>
          <Link to="/gallery">Gallery</Link>
          <Link to="/about">About</Link>
          <Link to="/contact">Contact</Link>
        </div>

        <div className="footer-column">
          <h3>Support</h3>
          <Link to="/faq">FAQs</Link>
          <Link to="/shipping">Shipping & Delivery</Link>
          <Link to="/returns">Returns</Link>
          <Link to="/privacy">Privacy Policy</Link>
          <Link to="/terms">Terms & Conditions</Link>
        </div>

        <div className="footer-column newsletter">
          <h3>Newsletter</h3>
          <p>Stay updated with new artworks and exhibition news.</p>

          <div className="newsletter-input">
            <input type="email" placeholder="Enter your email" />
            <button aria-label="Subscribe">→</button>
          </div>
        </div>

        <div className="copyright">
          © 2026 Nalini Art Gallery. All rights reserved.
        </div>
      </footer>
    </div>
  );
}

export default Home;