"""Unit tests for weather data conversion functions.

Comprehensive tests for all unit conversion functions including:
- Temperature (Fahrenheit to Celsius)
- Pressure (inHg to hPa)
- Speed (mph to km/h)
- Length (inches to mm)
- Wind direction (degrees to cardinal)
"""



class TestFahrenheitToCelsius:
    """Tests for Fahrenheit to Celsius conversion."""

    def test_freezing_point(self):
        from weatherstation import f_to_c

        assert f_to_c(32) == 0.0

    def test_boiling_point(self):
        from weatherstation import f_to_c

        assert f_to_c(212) == 100.0

    def test_body_temperature(self):
        from weatherstation import f_to_c

        assert f_to_c(98.6) == 37.0

    def test_negative_fahrenheit(self):
        from weatherstation import f_to_c

        assert f_to_c(-40) == -40.0  # -40 is same in both scales

    def test_string_input(self):
        from weatherstation import f_to_c

        assert f_to_c("68") == 20.0

    def test_room_temperature(self):
        from weatherstation import f_to_c

        # 70°F should be approximately 21.1°C
        result = f_to_c(70)
        assert 21.0 <= result <= 21.2


class TestInHgToHpa:
    """Tests for inches of mercury to hectopascals conversion."""

    def test_standard_pressure(self):
        from weatherstation import inhg_to_hpa

        # 29.92 inHg is standard atmospheric pressure (~1013 hPa)
        result = inhg_to_hpa(29.92)
        assert 1013.0 <= result <= 1014.0

    def test_low_pressure(self):
        from weatherstation import inhg_to_hpa

        result = inhg_to_hpa(29.0)
        assert 982.0 <= result <= 983.0

    def test_high_pressure(self):
        from weatherstation import inhg_to_hpa

        result = inhg_to_hpa(30.5)
        assert 1032.0 <= result <= 1034.0

    def test_string_input(self):
        from weatherstation import inhg_to_hpa

        result = inhg_to_hpa("29.92")
        assert 1013.0 <= result <= 1014.0

    def test_exact_conversion(self):
        from weatherstation import inhg_to_hpa

        # 29.92 inHg = 1013.2 hPa (rounded to 1 decimal)
        assert inhg_to_hpa(29.92) == 1013.2
        assert inhg_to_hpa(30.00) == 1015.9


class TestMphToKmh:
    """Tests for miles per hour to kilometers per hour conversion."""

    def test_zero(self):
        from weatherstation import mph_to_kmh

        assert mph_to_kmh(0) == 0.0

    def test_highway_speed(self):
        from weatherstation import mph_to_kmh

        # 60 mph ≈ 96.6 km/h
        result = mph_to_kmh(60)
        assert 96.5 <= result <= 96.6

    def test_string_input(self):
        from weatherstation import mph_to_kmh

        result = mph_to_kmh("10")
        assert 16.0 <= result <= 16.1

    def test_exact_conversion(self):
        from weatherstation import mph_to_kmh

        assert mph_to_kmh(10) == 16.1
        assert mph_to_kmh(60) == 96.6


class TestInchToMm:
    """Tests for inches to millimeters conversion."""

    def test_one_inch(self):
        from weatherstation import inch_to_mm

        assert inch_to_mm(1) == 25.4

    def test_zero(self):
        from weatherstation import inch_to_mm

        assert inch_to_mm(0) == 0.0

    def test_fractional(self):
        from weatherstation import inch_to_mm

        result = inch_to_mm(0.5)
        assert result == 12.7

    def test_string_input(self):
        from weatherstation import inch_to_mm

        assert inch_to_mm("2") == 50.8


class TestDegreesToCardinal:
    """Tests for degrees to cardinal direction conversion."""

    def test_primary_directions(self):
        """Test the four primary compass directions."""
        from weatherstation import degrees_to_cardinal

        assert degrees_to_cardinal(0) == "N"
        assert degrees_to_cardinal(90) == "E"
        assert degrees_to_cardinal(180) == "S"
        assert degrees_to_cardinal(270) == "W"

    def test_16_point_compass_rose(self):
        """Test all 16 compass directions."""
        from weatherstation import degrees_to_cardinal

        assert degrees_to_cardinal(0) == "N"
        assert degrees_to_cardinal(22.5) == "NNE"
        assert degrees_to_cardinal(45) == "NE"
        assert degrees_to_cardinal(67.5) == "ENE"
        assert degrees_to_cardinal(90) == "E"
        assert degrees_to_cardinal(112.5) == "ESE"
        assert degrees_to_cardinal(135) == "SE"
        assert degrees_to_cardinal(157.5) == "SSE"
        assert degrees_to_cardinal(180) == "S"
        assert degrees_to_cardinal(202.5) == "SSW"
        assert degrees_to_cardinal(225) == "SW"
        assert degrees_to_cardinal(247.5) == "WSW"
        assert degrees_to_cardinal(270) == "W"
        assert degrees_to_cardinal(292.5) == "WNW"
        assert degrees_to_cardinal(315) == "NW"
        assert degrees_to_cardinal(337.5) == "NNW"

    def test_360_degrees(self):
        """Test that 360° equals North."""
        from weatherstation import degrees_to_cardinal

        assert degrees_to_cardinal(360) == "N"

    def test_boundary_values(self):
        """Test degrees to cardinal conversion at boundary values."""
        from weatherstation import degrees_to_cardinal

        # Test boundaries between directions (offset by 11.25°)
        assert degrees_to_cardinal(11.24) == "N"  # Just before NNE
        assert degrees_to_cardinal(11.25) == "NNE"  # Exactly at boundary
        assert degrees_to_cardinal(33.74) == "NNE"  # Just before NE
        assert degrees_to_cardinal(33.75) == "NE"  # Exactly at boundary

    def test_wrap_around(self):
        """Test degrees to cardinal handles values > 360."""
        from weatherstation import degrees_to_cardinal

        # Should normalize to 0-360 range
        assert degrees_to_cardinal(405) == "NE"  # 405 % 360 = 45
        assert degrees_to_cardinal(720) == "N"  # 720 % 360 = 0
        assert degrees_to_cardinal(450) == "E"  # 450 % 360 = 90

    def test_none_input(self):
        """Test degrees to cardinal handles None input."""
        from weatherstation import degrees_to_cardinal

        assert degrees_to_cardinal(None) is None

    def test_invalid_input(self):
        """Test degrees to cardinal handles invalid input."""
        from weatherstation import degrees_to_cardinal

        assert degrees_to_cardinal("invalid") is None
        assert degrees_to_cardinal("") is None

    def test_string_numeric_input(self):
        """Test degrees to cardinal handles string numeric input."""
        from weatherstation import degrees_to_cardinal

        assert degrees_to_cardinal("180") == "S"
        assert degrees_to_cardinal("45.0") == "NE"
