﻿using Explorer.Stakeholders.API.Dtos;
using Explorer.Stakeholders.Core.Domain;
using Explorer.Stakeholders.Core.UseCases;
using FluentResults;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;

namespace Explorer.Stakeholders.Infrastructure.Authentication;

public class JwtGenerator : ITokenGenerator
{
    private readonly string _key = Environment.GetEnvironmentVariable("JWT_KEY") ?? "explorer_secret_key";
    private readonly string _issuer = Environment.GetEnvironmentVariable("JWT_ISSUER") ?? "explorer";
    private readonly string _audience = Environment.GetEnvironmentVariable("JWT_AUDIENCE") ?? "explorer-front.com";

    public Result<AuthenticationTokensDto> GenerateAccessToken(User user, long personId)
    {
        var authenticationResponse = new AuthenticationTokensDto();

        var claims = new List<Claim>
        {
            new(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
            new("id", user.Id.ToString()),
            new("username", user.Username),
            new("personId", personId.ToString()),
            new(ClaimTypes.Role, user.GetPrimaryRoleName())
        };

        var jwt = CreateToken(claims, 60 * 24);
        authenticationResponse.Id = user.Id;
        authenticationResponse.AccessToken = jwt;

        return authenticationResponse;
    }

    private string CreateToken(IEnumerable<Claim> claims, double expirationTimeInMinutes)
    {
        var securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_key));
        var credentials = new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256);

        var token = new JwtSecurityToken(
            _issuer,
            _audience,
            claims,
            expires: DateTime.Now.AddMinutes(expirationTimeInMinutes),
            signingCredentials: credentials);

        return new JwtSecurityTokenHandler().WriteToken(token);
    }


    public string GenerateResetPasswordToken(User user, long personId)
    {
        var authenticationResponse = new AuthenticationTokensDto();

        var claims = new List<Claim>
        {
            new(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
            new("id", user.Id.ToString()),
            new("username", user.Username),
            new("personId", personId.ToString()),
            new(ClaimTypes.Role, user.GetPrimaryRoleName())
        };

        var jwt = CreateToken(claims, 60 * 24);

        return jwt;
    }

    public string GenerateEmailVerificationToken(string email, string username)
    {
        var authenticationResponse = new AuthenticationTokensDto();

        var claims = new List<Claim>
        {
            new(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
            new("username", username),
            new("email", email)
        };

        var jwt = CreateToken(claims, 60 * 24);


        return jwt;
    }


    public long GetUserIdFromToken(string jwtToken)
    {
        var handler = new JwtSecurityTokenHandler();
        var jsonToken = handler.ReadToken(jwtToken) as JwtSecurityToken;

        if (jsonToken?.Payload != null && jsonToken.Payload.TryGetValue("id", out var userId))
        {
            if (userId is string userIdString && long.TryParse(userIdString, out var userIdLong))
            {
                return userIdLong;
            }
        }

        return 0;
    }

    public string GetUserEmailFromToken(string jwtToken)
    {
        var handler = new JwtSecurityTokenHandler();
        var jsonToken = handler.ReadToken(jwtToken) as JwtSecurityToken;

        if (jsonToken?.Payload != null && jsonToken.Payload.TryGetValue("email", out var userEmail))
        {
            if (userEmail is string userEmailString)
            {
                return userEmailString;
            }
        }

        return null;
    }

    public DateTime GetTokenExpirationTime(string jwtToken)
    {
        var handler = new JwtSecurityTokenHandler();
        var jsonToken = handler.ReadToken(jwtToken) as JwtSecurityToken;

        if (jsonToken?.Payload != null && jsonToken.Payload.TryGetValue("exp", out var expirationTime))
        {
            if (expirationTime is long expirationTimeUnix)
            {
                return UnixTimeStampToDateTime(expirationTimeUnix);
            }
        }

        return DateTime.UtcNow;
    }

    private DateTime UnixTimeStampToDateTime(long unixTimeStamp)
    {
        var dateTimeOffset = DateTimeOffset.FromUnixTimeSeconds(unixTimeStamp);
        return dateTimeOffset.UtcDateTime;
    }
}