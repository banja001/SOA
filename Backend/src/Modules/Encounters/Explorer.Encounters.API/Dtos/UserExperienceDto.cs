using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace Explorer.Encounters.API.Dtos
{
    public class UserExperienceDto
    {
        [JsonPropertyName("id")]
        public long Id { get; set; }

        [JsonPropertyName("userId")]
        public long UserId { get; set; }

        [JsonPropertyName("xp")]
        public int XP { get; set; }

        [JsonPropertyName("level")]
        public int Level { get; set; }
        public UserExperienceDto(long id, long userId, int xp, int level)
        {
            Id = id;
            UserId = userId;
            XP = xp;
            Level = level;
        }
        public UserExperienceDto()
        {
        }
        public UserExperienceDto(long userId, int xP, int level)
        {
            UserId = userId;
            XP = xP;
            Level = level;
        }
    }
}
