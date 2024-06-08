# URLShortner

This project aims to develop a URL shortener service that allows users to convert long URLs into shorter, more manageable strings. These shortened URLs can then be easily shared and used in various contexts, such as social media posts, emails, or printed materials.

## Features
- URL Shortening: Users can submit long URLs, and the service will generate a unique, shorter representation of the original URL.
- Custom Short Codes: Users might be able to create custom short codes for their shortened URLs (subject to availability and validation rules).
- URL Expiration: Implement an option for URLs to expire after a certain period, ensuring they don't point to outdated content.
- Click Tracking: Track the number of clicks on each shortened URL for analytics purposes.
- User Accounts: Implement user accounts to manage shortened URLs and access click tracking data. 
- Advanced Analytics: Provide detailed click tracking reports with user demographics and geographical data. 

## Technology Stacks

- Programming Language: Golang
- Database: MySQL
- URL Shortening Algorithm: 

## Benefits

- Improved Link Management: Shortened URLs are easier to share and manage compared to lengthy original URLs. 
- Click Tracking: Provides valuable insights into user behavior and website traffic patterns. 
- Customizable Short Codes: Enhances user experience by allowing personalized short URLs.

## Key Considerations

### Unique ID Generations:
- Collision Resistance: Ensure that the each shortened url is unique.
- Length of shortened urls: A middle ground between long and short urls should be reached to avoid collisions.
- Base Encoding: Use base64 encoding to convert IDs.

### Database Design:
- Table Schema: Original url to shortened url mappings.
- Indexing: Index shortened urls column for faster lookups.

### Redirection logic:
- Implement logic to handle the redirection from the short url to the original.
- Handle edge cases like invalid or expired urls.