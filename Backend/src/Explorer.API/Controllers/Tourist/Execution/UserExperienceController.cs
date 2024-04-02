using System.Text.Json;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Encounters.API.Dtos;
using Explorer.Encounters.API.Public;
using Explorer.Encounters.Core.UseCases;
using Explorer.Tours.Core.Domain.Tours;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Tourist.Execution
{
    [Authorize(Policy = "touristPolicy")]
    [Route("api/tourist/userExperience")]
    public class UserExperienceController : BaseApiController
    {
        private readonly IUserExperienceService _userExperienceService;
        private readonly IChallengeExecutionService _challengeExecutionService;
        private readonly IHttpClientFactory _factory;

        public UserExperienceController(IUserExperienceService userExperienceService, IChallengeExecutionService challengeExecutionService, IHttpClientFactory factory)
        {
            _userExperienceService = userExperienceService;
            _challengeExecutionService = challengeExecutionService;
            _factory = factory;
        }
        [HttpGet]
        public ActionResult<PagedResult<UserExperienceDto>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _userExperienceService.GetPaged(page, pageSize);
            return CreateResponse(result);
        }

        [HttpPost]
        public async Task<ActionResult<UserExperienceDto>> Create([FromBody] UserExperienceDto userExperience)
        {
            //var result = _userExperienceService.Create(userExperience);
            //return CreateResponse(result);
            var client = _factory.CreateClient("encountersMicroservice");

            using HttpResponseMessage response = await client.PostAsJsonAsync("/userxp/create", userExperience);

            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();

            return Ok(jsonResponse);
        }

        [HttpPut]
        public async Task<ActionResult<UserExperienceDto>> Update([FromBody] UserExperienceDto userExperience)
        {
            //var result = _userExperienceService.Update(userExperience);
            //return CreateResponse(result);
            var client = _factory.CreateClient("encountersMicroservice");
            
            using HttpResponseMessage response = await client.PutAsJsonAsync("/userxp/update", userExperience);

            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();

            return Ok(jsonResponse);
        }

        [HttpDelete("{id:int}")]
        public async Task<ActionResult> Delete(int id)
        {
            //var result = _userExperienceService.Delete(id);
            //return CreateResponse(result);
            var client = _factory.CreateClient("encountersMicroservice");

            using HttpResponseMessage response = await client.DeleteAsync("userxp/delete/" + id.ToString());
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = response.Content.ReadAsStringAsync();

            return Ok(jsonResponse);
        }

        [HttpGet("{userId:int}")]
        public async Task<ActionResult<PagedResult<UserExperienceDto>>> GetByUserId(int userId)
        {
            //var result = _userExperienceService.GetByUserId(userId);
            var client = _factory.CreateClient("encountersMicroservice");

            using HttpResponseMessage response = await client.GetAsync("userxp/" + userId.ToString());
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }
            
            var jsonResponse = await response.Content.ReadAsStringAsync();
            
            var userExperienceDto = JsonSerializer.Deserialize<UserExperienceDto>(jsonResponse);
            
            return Ok(userExperienceDto);
        }

        [HttpPut("addxp/{id:int}/{xp:int}")]
        public async Task<ActionResult> AddXP(int id,int xp)
        {
            //var result = _userExperienceService.AddXP(id,xp);
            var client = _factory.CreateClient("encountersMicroservice");

            using HttpResponseMessage response = await client.PutAsync("userxp/add/" + id.ToString() + "/" + xp.ToString(),  null);
        
            if (!response.IsSuccessStatusCode)
            {
                return StatusCode((int)response.StatusCode);
            }

            var jsonResponse = await response.Content.ReadAsStringAsync();

            return Ok(jsonResponse);
        }

        [HttpPut("addxpsocial/{challengeId:long}/{xp:int}")]
        public ActionResult<UserExperienceDto> AddXPSocial(long challengeId, int xp)
        {
            var ids= _challengeExecutionService.GetUserIds(challengeId);
            Result<UserExperienceDto> result=new Result<UserExperienceDto>();
            foreach (var id in ids.Value) 
            {
                result = _userExperienceService.AddXP(_userExperienceService.GetByUserId(id).Value.Id, xp);
            }
            return CreateResponse(result);
        }
    }
}
