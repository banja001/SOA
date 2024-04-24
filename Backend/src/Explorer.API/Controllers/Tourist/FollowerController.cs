using Explorer.Encounters.API.Dtos;
using Explorer.Stakeholders.API.Dtos;
using Explorer.Stakeholders.API.Public.Identity;
using Explorer.Stakeholders.Core.Domain;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Text.Json;

namespace Explorer.API.Controllers.Tourist.Identity
{
    [Authorize(Policy = "touristPolicy")]
    [Route("api/tourist/follower")]
    public class FollowerController : BaseApiController
    {
        private readonly IFollowerService _followerService;
        private readonly IHttpClientFactory _factory;

        public FollowerController(IFollowerService followerService, IHttpClientFactory factory)
        {
            _followerService = followerService;
            _factory= factory;
        }

        [HttpGet("{id:int}")]
        public ActionResult<List<FollowerDto>> GetFollowersNotifications(int id)
        {
            var result = _followerService.GetFollowersNotifications(id);
            return CreateResponse(result);
        }

        [HttpPut]
        public async Task<ActionResult<FollowerDto>> Create([FromBody] FollowerDto follower)
        {
            var client = _factory.CreateClient("followerMicroservice");
            using HttpResponseMessage response = await client.PutAsJsonAsync("followers/update", follower);
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();

            return Ok(jsonResponse);
        }

        [HttpDelete("{followerId:int}/{followedId:int}")]
        public ActionResult Delete(int followerId, int followedId)
        {
            var result = _followerService.Delete(followerId, followedId);
            return CreateResponse(result);
        }

        [HttpGet("followings/{id:int}/{uid:int}")]
        public async Task<ActionResult<List<PersonGoDto>>> GetFollowings(int id, int uid)
        {
            var client = _factory.CreateClient("followerMicroservice");

            using HttpResponseMessage response = await client.GetAsync("followers/recommended/" + id.ToString()+"/"+ uid.ToString());
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();

            List<PersonGoDto> persons = JsonSerializer.Deserialize<List<PersonGoDto>>(jsonResponse);

            return Ok(persons);
        }
    }
}
