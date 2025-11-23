"""Tests for Weather Underground forwarding functionality.

Tests for optional WU data forwarding including:
- Forwarding when enabled
- No forwarding when disabled
- Error handling
"""

from unittest.mock import Mock, patch


class TestWeatherUndergroundForwarding:
    """Test Weather Underground forwarding functionality."""

    @patch("weatherstation.WU_FORWARD", True)
    @patch("weatherstation.WU_USERNAME", "test_user")
    @patch("weatherstation.WU_PASSWORD", "test_pass")
    @patch("weatherstation.requests.get")
    @patch("weatherstation.dns.resolver.Resolver")
    def test_wu_forwarding_enabled(self, mock_resolver, mock_requests, client, mock_mqtt):
        """Test data forwarded to Weather Underground when enabled."""
        # Mock DNS resolution
        mock_answer = Mock()
        mock_answer.to_text.return_value = "1.2.3.4"
        mock_resolver.return_value.resolve.return_value = [mock_answer]

        # Mock HTTP request
        mock_response = Mock()
        mock_response.status_code = 200
        mock_requests.return_value = mock_response

        response = client.get("/weatherstation/updateweatherstation.php?tempf=70")

        assert response.status_code == 200
        assert mock_requests.called

    @patch("weatherstation.WU_FORWARD", False)
    @patch("weatherstation.requests.get")
    def test_wu_forwarding_disabled(self, mock_requests, client, mock_mqtt):
        """Test data not forwarded when WU forwarding disabled."""
        response = client.get("/weatherstation/updateweatherstation.php?tempf=70")

        assert response.status_code == 200
        assert not mock_requests.called

    @patch("weatherstation.WU_FORWARD", True)
    @patch("weatherstation.WU_USERNAME", "test_user")
    @patch("weatherstation.WU_PASSWORD", "test_pass")
    @patch("weatherstation.requests.get")
    @patch("weatherstation.dns.resolver.Resolver")
    def test_wu_forwarding_error_does_not_affect_response(
        self, mock_resolver, mock_requests, client, mock_mqtt
    ):
        """Test that WU forwarding errors don't affect the main response."""
        # Mock DNS resolution to raise an exception
        mock_resolver.return_value.resolve.side_effect = Exception("DNS error")

        response = client.get("/weatherstation/updateweatherstation.php?tempf=70")

        # Should still return success to weather station
        assert response.status_code == 200
        assert response.data == b"success"

    @patch("weatherstation.WU_FORWARD", True)
    @patch("weatherstation.WU_USERNAME", "test_user")
    @patch("weatherstation.WU_PASSWORD", "test_pass")
    @patch("weatherstation.requests.get")
    @patch("weatherstation.dns.resolver.Resolver")
    def test_wu_forwarding_http_error_handled(
        self, mock_resolver, mock_requests, client, mock_mqtt
    ):
        """Test that HTTP errors during WU forwarding are handled gracefully."""
        # Mock DNS resolution
        mock_answer = Mock()
        mock_answer.to_text.return_value = "1.2.3.4"
        mock_resolver.return_value.resolve.return_value = [mock_answer]

        # Mock HTTP request to fail
        mock_requests.side_effect = Exception("Connection timeout")

        response = client.get("/weatherstation/updateweatherstation.php?tempf=70")

        # Should still return success to weather station
        assert response.status_code == 200
        assert response.data == b"success"

    @patch("weatherstation.WU_FORWARD", True)
    @patch("weatherstation.WU_USERNAME", "")
    @patch("weatherstation.WU_PASSWORD", "")
    @patch("weatherstation.requests.get")
    def test_wu_forwarding_without_credentials(self, mock_requests, client, mock_mqtt):
        """Test WU forwarding behavior when credentials are empty."""
        response = client.get("/weatherstation/updateweatherstation.php?tempf=70")

        # Should still return success
        assert response.status_code == 200
